package main

import (
	"fmt"
	"net/url"
	"slices"
)

func CleanYouTubeURL(inputURL string) (string, error) {
	hosts := []string{
		"youtu.be",
		"youtube.com",
	}
	u, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %w", err)
	}
	if !slices.Contains(hosts, u.Host) {
		return "", ErrNotYouTube
	}
	params := u.Query()
	if !params.Has("si") {
		return inputURL, nil
	}
	params.Del("si")
	u.RawQuery = params.Encode()
	return u.String(), nil
}
