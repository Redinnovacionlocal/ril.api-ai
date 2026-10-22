package guardrails_plugin

import "google.golang.org/adk/session"

// Input is what each rule checks.
type Input struct {
	UserText     string
	ResponseText string
	State        session.State
}

// Violation describes a failed rule.
type Violation struct {
	Rule string
	// Detail is logged.
	Detail string
	// Instruction is sent to the rewrite model.
	Instruction string
	// Fragments limits the rewrite to these exact pieces of the response.
	// Empty means the whole response.
	Fragments []string
	// OnKeep runs when the rewrite model keeps the response because the user
	// explicitly asked for it. Optional.
	OnKeep func()
}

// Rule validates a final response. Check runs on every response, so keep it
// cheap. Returns nil if the response passes.
type Rule interface {
	Name() string
	Check(in Input) *Violation
}

// LineRule is a Rule that can also check a streamed response line by line,
// before each line reaches the user.
type LineRule interface {
	Rule
	// LineChecker returns nil when the rule does not apply to this response.
	LineChecker(in Input) func(line string) *Violation
}

// ruleFactory lazily builds a rule, only when enabled.
type ruleFactory struct {
	name string
	new  func() (Rule, error)
}

// registeredRules lists all available rules.
func registeredRules() []ruleFactory {
	return []ruleFactory{
		{name: "language", new: newLanguageRule},
	}
}
