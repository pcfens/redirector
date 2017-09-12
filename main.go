package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/pcfens/redirector/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	var configFile = flag.String("config", "config.yaml", "Location of the configuration file to use")
	flag.Parse()
	// TODO: Allow config variables to be set at the command line
	configuration, err := config.LoadConfig(*configFile)

	if err != nil {
		log.Panic("Error reading configuration")
	}

	http.HandleFunc("/", configuration.RedirectList.HTTPHandler)
	http.HandleFunc("/reload", configuration.HTTPHandler)

	log.Info("Binding HTTP server to ", configuration.BindAddress)
	panic(http.ListenAndServe(configuration.BindAddress, nil))
}
