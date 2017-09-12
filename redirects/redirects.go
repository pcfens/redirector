package redirects

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// Redirects is a slice of redirects
type Redirects []Redirect

// Redirect is an individual redirect with original request and potential
// redirect targets.
type Redirect struct {
	Name    string `yaml:"name"`
	Pattern string `yaml:"pattern"`
	re      *regexp.Regexp
	Targets Targets `yaml:"targets"`
}

func (redirects Redirects) combine(r Redirects) Redirects {
	for _, redirect := range r {
		redirects = append(redirects, redirect)
	}
	return redirects
}

// FindDest will find where the user should be redirected to
func (redirects Redirects) FindDest(r *http.Request) Target {
	url := fmt.Sprintf("%v%v\n", r.Host, r.RequestURI)

	for _, option := range redirects {
		if option.re.MatchString(url) {
			dest := option.Targets.findTarget(r)
			dest.Target = strings.TrimSpace(option.re.ReplaceAllString(url, dest.Target))
			return dest
		}
	}
	return (Target{
		Code: http.StatusNotFound,
	})
}
