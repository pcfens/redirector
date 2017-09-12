package redirects

import (
	"net"
	"net/http"
	"strings"
)

// Condition is a set of conditions that must be met for the target to be used
type Condition struct {
	Source []string `yaml:"source,omitempty"`
}

// IsEmpty determines whether or not there are conditions
func (c Condition) IsEmpty() bool {
	return len(c.Source) == 0
}

// DoesConditionMatch returns true if all conditions are met in the supplied
// http.Request object.
func (c Condition) DoesConditionMatch(r *http.Request) bool {
	if c.IsEmpty() {
		return true
	}

	for _, cidr := range c.Source {
		ips := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
		_, network, err := net.ParseCIDR(cidr)
		if err != nil {
			return false
		}
		for _, ip := range ips {
			if network.Contains(net.ParseIP(ip)) {
				return true
			}
		}
	}
	return false
}
