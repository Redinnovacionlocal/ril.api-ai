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
	minWords      = 3
	minConfidence = 0.7
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

func (r *languageRule) Check(in Input) *Violation {
	msgLang, msgConf := r.detect(in.UserText)
	storedLang := r.storedLanguage(in.State)
	userLang := msgLang
	if userLang == lingua.Unknown {
		userLang = storedLang
	} else if userLang != storedLang {
		r.storeLanguage(in.State, userLang)
	}
	respLang, respConf := r.detect(in.ResponseText)

	slog.Debug("guardrails: language check",
		"message_lang", isoCode(msgLang), "message_confidence", msgConf,
		"stored_lang", isoCode(storedLang), "user_lang", isoCode(userLang),
		"response_lang", isoCode(respLang), "response_confidence", respConf)

	if userLang == lingua.Unknown || respLang == lingua.Unknown || respLang == userLang {
		return nil
	}
	return &Violation{
		Rule:   r.Name(),
		Detail: fmt.Sprintf("user=%s response=%s", isoCode(userLang), isoCode(respLang)),
		Instruction: fmt.Sprintf(
			"La respuesta está en %s pero el usuario escribe en %s: hay que traducirla completa a %s.",
			respLang, userLang, userLang),
		OnKeep: func() { r.storeLanguage(in.State, respLang) },
	}
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
