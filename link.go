package main

import "regexp"

var urlRegex = regexp.MustCompile(`https?://[^\s"'<>]+`)

func ExtractURL(text string) (string, error) {
	match := urlRegex.FindString(text)

	if match == "" {
		return "", ErrNoURL
	}

	return match, nil
}
