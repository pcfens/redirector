package redirects

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HTTPHandler sends an HTTP redirect based on the request
func (redirects *Redirects) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	dest := redirects.FindDest(r)
	log.WithFields(log.Fields{
		"target": dest.Target,
		"code":   dest.Code,
		"client": r.RemoteAddr,
	}).Info(r.Host, r.RequestURI)

	if dest.Code == http.StatusNotFound {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, dest.Target, dest.Code)
	}
}
