package markdown

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
)

func ReadURL() (*string, error) {
	uri := flag.String("uri", "", "URL of the markdown file to parse. (required)")
	flag.Parse()

	if *uri == "" {
		return nil, fmt.Errorf("uri is required")
	}

	if u, err := url.Parse(*uri); err != nil || u.Scheme != "https" || u.Host != "raw.githubusercontent.com" {
		return nil, fmt.Errorf("invalid uri: %w", err)
	}

	response, err := http.Head(*uri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch the markdown file cntent: status %d", response.StatusCode)
	}

	return uri, nil
}
