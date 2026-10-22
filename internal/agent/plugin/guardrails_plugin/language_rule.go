package guardrails_plugin

import (
	"errors"
	"fmt"
	"log/slog"
	"regexp"
	"strings"
	"sync"

	"github.com/pemistahl/lingua-go"
	"google.golang.org/adk/session"
	"ril.api-ia/internal/infrastructure/env"
)

const userLanguageStateKey = "guardrails_user_lang"

const (
	minWords          = 3
	minParagraphWords = 5
	minConfidence     = 0.7
)

var (
	codeBlockRe = regexp.MustCompile("(?s)```.*?```|`[^`]*`")
	urlRe       = regexp.MustCompile(`https?://\S+|www\.\S+`)

	detectorOnce sync.Once
	detector     lingua.LanguageDetector
	detectorErr  error
)

type languageRule struct {
	detector lingua.LanguageDetector
}

func newLanguageRule() (Rule, error) {
	detectorOnce.Do(func() {
		detector, detectorErr = buildDetector(env.GetOrDefault("GUARDRAILS_LANGUAGES", "es,pt,en"))
	})
	if detectorErr != nil {
		return nil, detectorErr
	}
	return &languageRule{detector: detector}, nil
}

// buildDetector restricts detection to the configured languages.
func buildDetector(codes string) (lingua.LanguageDetector, error) {
	var languages []lingua.Language
	for _, code := range strings.Split(codes, ",") {
		code = strings.TrimSpace(code)
		if code == "" {
			continue
		}
		lang := lingua.GetLanguageFromIsoCode639_1(lingua.GetIsoCode639_1FromValue(code))
		if lang == lingua.Unknown {
			return nil, fmt.Errorf("GUARDRAILS_LANGUAGES: unknown language %q", code)
		}
		languages = append(languages, lang)
	}
	if len(languages) < 2 {
		return nil, errors.New("GUARDRAILS_LANGUAGES: at least 2 languages are required")
	}
	return lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		WithPreloadedLanguageModels().
		Build(), nil
}

func (r *languageRule) Name() string {
	return "language"
}

// Check accepts the conversation language (stored in the session) or the
// current message language. A single message does not change the conversation
// language: it changes only when the response follows the message.
func (r *languageRule) Check(in Input) *Violation {
	msgLang, msgConf := r.detect(in.UserText)
	convLang := r.storedLanguage(in.State)
	if convLang == lingua.Unknown && msgLang != lingua.Unknown {
		convLang = msgLang
		r.storeLanguage(in.State, convLang)
	}
	respLang, respConf := r.detect(in.ResponseText)
	foreign := r.foreignParagraphs(in.ResponseText, convLang, msgLang)

	slog.Debug("guardrails: language check",
		"message_lang", isoCode(msgLang), "message_confidence", msgConf,
		"conversation_lang", isoCode(convLang),
		"response_lang", isoCode(respLang), "response_confidence", respConf,
		"foreign_paragraphs", len(foreign))

	if convLang == lingua.Unknown {
		return nil
	}
	expected := func(l lingua.Language) bool { return l == convLang || l == msgLang }

	if respLang != lingua.Unknown && !expected(respLang) {
		return &Violation{
			Rule:   r.Name(),
			Detail: fmt.Sprintf("conversation=%s message=%s response=%s", isoCode(convLang), isoCode(msgLang), isoCode(respLang)),
			Instruction: fmt.Sprintf(
				"La respuesta está en %s pero la conversación es en %s: hay que traducirla completa a %s.",
				respLang, convLang, convLang),
			OnKeep: func() { r.storeLanguage(in.State, respLang) },
		}
	}
	if len(foreign) > 0 {
		return r.foreignViolation(convLang, foreign)
	}
	if respLang == msgLang && msgLang != convLang {
		r.storeLanguage(in.State, msgLang)
	}
	return nil
}

// LineChecker flags English lines while the response is streamed.
func (r *languageRule) LineChecker(in Input) func(line string) *Violation {
	msgLang, _ := r.detect(in.UserText)
	convLang := r.storedLanguage(in.State)
	if convLang == lingua.Unknown {
		convLang = msgLang
	}
	if !r.checksLines(convLang, msgLang) {
		return nil
	}
	return func(line string) *Violation {
		if !r.isEnglishLine(line) {
			return nil
		}
		return r.foreignViolation(convLang, []string{line})
	}
}

func (r *languageRule) foreignViolation(convLang lingua.Language, lines []string) *Violation {
	return &Violation{
		Rule:        r.Name(),
		Detail:      fmt.Sprintf("conversation=%s paragraphs=%d", isoCode(convLang), len(lines)),
		Instruction: fmt.Sprintf("Hay fragmentos en inglés pero la conversación es en %s.", convLang),
		Fragments:   lines,
	}
}

// foreignParagraphs returns the lines in English when the conversation is
// not, usually leaked reasoning: Gemini reasons in English. Code blocks are
// skipped.
func (r *languageRule) foreignParagraphs(text string, convLang, msgLang lingua.Language) []string {
	if !r.checksLines(convLang, msgLang) {
		return nil
	}
	var foreign []string
	inCode := false
	for _, line := range strings.Split(text, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inCode = !inCode
			continue
		}
		if !inCode && r.isEnglishLine(line) {
			foreign = append(foreign, line)
		}
	}
	return foreign
}

func (r *languageRule) checksLines(convLang, msgLang lingua.Language) bool {
	return convLang != lingua.Unknown && convLang != lingua.English && msgLang != lingua.English
}

// isEnglishLine checks only for English because lingua mixes up Spanish and
// Portuguese in short texts. Short lines are skipped.
func (r *languageRule) isEnglishLine(line string) bool {
	if len(strings.Fields(line)) < minParagraphWords {
		return false
	}
	lang, _ := r.detect(line)
	return lang == lingua.English
}

func (r *languageRule) storedLanguage(state session.State) lingua.Language {
	if state == nil {
		return lingua.Unknown
	}
	v, err := state.Get(userLanguageStateKey)
	if err != nil {
		if !errors.Is(err, session.ErrStateKeyNotExist) {
			slog.Warn("guardrails: error reading language from session", "error", err)
		}
		return lingua.Unknown
	}
	code, _ := v.(string)
	return lingua.GetLanguageFromIsoCode639_1(lingua.GetIsoCode639_1FromValue(code))
}

func (r *languageRule) storeLanguage(state session.State, lang lingua.Language) {
	if state == nil {
		return
	}
	if err := state.Set(userLanguageStateKey, isoCode(lang)); err != nil {
		slog.Warn("guardrails: error saving language to session", "error", err)
	}
}

// detect returns lingua.Unknown when the text is too short or ambiguous.
// The confidence is returned anyway for logging.
func (r *languageRule) detect(text string) (lingua.Language, float64) {
	text = codeBlockRe.ReplaceAllString(text, " ")
	text = urlRe.ReplaceAllString(text, " ")
	if len(strings.Fields(text)) < minWords {
		return lingua.Unknown, 0
	}
	values := r.detector.ComputeLanguageConfidenceValues(text)
	if len(values) == 0 {
		return lingua.Unknown, 0
	}
	if values[0].Value() < minConfidence {
		return lingua.Unknown, values[0].Value()
	}
	return values[0].Language(), values[0].Value()
}

func isoCode(lang lingua.Language) string {
	if lang == lingua.Unknown {
		return "unknown"
	}
	return strings.ToLower(lang.IsoCode639_1().String())
}
