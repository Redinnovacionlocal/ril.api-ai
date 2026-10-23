package guardrails_plugin

import (
	"fmt"
	"regexp"
	"strings"
)

// latexRe matches LaTeX math: $$...$$, \(...\), \[...\], inline $...\cmd...$
// and common commands.
var latexRe = regexp.MustCompile(`\$\$|\\\(|\\\[|\$[^$\n]*\\[a-zA-Z]+[^$\n]*\$|\\(text|frac|times|cdot|sum|sqrt|left|right|mathrm|leq|geq|approx)\b`)

type latexRule struct{}

func newLatexRule() (Rule, error) {
	return latexRule{}, nil
}

func (latexRule) Name() string {
	return "latex"
}

// Check flags the lines with LaTeX math, which the chat does not render.
// Code blocks are skipped.
func (r latexRule) Check(in Input) *Violation {
	var lines []string
	inCode := false
	for _, line := range strings.Split(in.ResponseText, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inCode = !inCode
			continue
		}
		if !inCode && latexRe.MatchString(line) {
			lines = append(lines, line)
		}
	}
	if len(lines) == 0 {
		return nil
	}
	return r.violation(lines)
}

func (r latexRule) LineChecker(Input) func(line string) *Violation {
	return func(line string) *Violation {
		if !latexRe.MatchString(line) {
			return nil
		}
		return r.violation([]string{line})
	}
}

func (r latexRule) violation(lines []string) *Violation {
	return &Violation{
		Rule:   r.Name(),
		Detail: fmt.Sprintf("lines=%d", len(lines)),
		Instruction: "Hay fórmulas en formato LaTeX, que el chat no puede mostrar. Escribilas en texto plano " +
			"con palabras y símbolos comunes, sin $, \\text ni otros comandos. " +
			"Ejemplo: IPPI = Densidad de 0 a 5 años + Vulnerabilidad social − Cobertura de servicios.",
		Fragments: lines,
	}
}
