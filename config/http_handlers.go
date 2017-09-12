package config

import (
	"net/http"

	"github.com/pcfens/redirector/redirects"
	log "github.com/sirupsen/logrus"
)

// HTTPHandler reloads the redirect list
func (config *Config) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	log.WithFields(log.Fields{
		"client": r.RemoteAddr,
	}).Info("Reloading redirect list")

	config.RedirectList, err = redirects.LoadRedirects(config.RedirectSource)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Redirects reloaded"))
}
