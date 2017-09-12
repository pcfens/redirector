package redirects

import "net/http"

// Targets is a slice of multiptle Target objects
type Targets []Target

// Target is a potential redirect target with conditionals
type Target struct {
	Target string    `yaml:"target"`
	Code   int       `yaml:"code,omitempty"`
	When   Condition `yaml:"when,omitempty"`
}

// Determine which target is the most appropriate
func (targets Targets) findTarget(r *http.Request) Target {
	for _, candidateTarget := range targets {
		if candidateTarget.When.DoesConditionMatch(r) {
			return candidateTarget
		}
	}
	return Target{
		Code: http.StatusNotFound,
	}
}
