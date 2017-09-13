package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pcfens/redirector/redirects"

	yaml "gopkg.in/yaml.v2"
)

// Config is the configuration used at startup
type Config struct {
	BindAddress    string   `yaml:"bind_address"`
	IPHeader       string   `yaml:"ip_header"`
	RedirectSource []string `yaml:"redirects,omitempty"`
	RedirectList   redirects.Redirects
}

// LoadConfig reads the configuration and returns an instance of Config
func LoadConfig(configFile string) (Config, error) {
	var configuration = Config{
		BindAddress:    "0.0.0.0:8080",
		IPHeader:       "X-Forwareded-For",
		RedirectSource: []string{"rules.yaml"},
	}
	source, err := ioutil.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(source, &configuration)
	if err != nil {
		return Config{}, err
	}

	if os.Getenv("BIND_ADDRESS") != "" {
		configuration.BindAddress = os.Getenv("BIND_ADDRESS")
	}

	if os.Getenv("IP_HEADER") != "" {
		configuration.IPHeader = os.Getenv("IP_HEADER")
	}

	if os.Getenv("REDIRECT_SOURCE") != "" {
		configuration.RedirectSource = strings.Split(os.Getenv("REDIRECT_SOURCE"), ",")
	}

	configuration.RedirectList, err = redirects.LoadRedirects(configuration.RedirectSource)
	if err != nil {
		return Config{}, err
	}

	return configuration, nil
}
