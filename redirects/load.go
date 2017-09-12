package redirects

import (
	"path/filepath"
	"regexp"

	cloudfile "github.com/alexflint/go-cloudfile"

	yaml "gopkg.in/yaml.v2"
)

// LoadRedirects will load redirects from the configured source
func LoadRedirects(redirectSource []string) (Redirects, error) {
	//TODO: Implement loaders from git and/or URLs
	redirects, err := loadRedirectList(redirectSource)
	if err != nil {
		return Redirects{}, err
	}
	for index, redirect := range redirects {
		redirects[index].re = regexp.MustCompile(redirect.Pattern)
	}
	return redirects, nil
}

// Load redirects from an array of blobs
func loadRedirectList(redirectSource []string) (Redirects, error) {
	var redirects Redirects
	for _, entry := range redirectSource {
		files, err := filepath.Glob(entry)
		if err != nil {
			return Redirects{}, err
		}
		newRedirects := Redirects{}
		if len(files) > 0 {
			newRedirects, err = loadRedirectBlob(entry)
		} else {
			newRedirects, err = loadFromFile(entry)
		}
		if err != nil {
			return Redirects{}, err
		}
		redirects = redirects.combine(newRedirects)
	}
	return redirects, nil
}

// Load redirects from a Blob
func loadRedirectBlob(blob string) (Redirects, error) {
	var redirects Redirects
	files, err := filepath.Glob(blob)
	if err != nil {
		return Redirects{}, err
	}
	for _, file := range files {
		newRedirects, err := loadFromFile(file)
		if err != nil {
			return Redirects{}, err
		}
		redirects = redirects.combine(newRedirects)
	}
	return redirects, nil
}

// Load redirects from an individual file
func loadFromFile(filename string) (Redirects, error) {
	var redirects Redirects
	source, err := cloudfile.ReadFile(filename)
	// source, err := ioutil.ReadFile(filename)
	if err != nil {
		return Redirects{}, err
	}
	err = yaml.Unmarshal(source, &redirects)
	if err != nil {
		return Redirects{}, err
	}

	return redirects, nil
}
