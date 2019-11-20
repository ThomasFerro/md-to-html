package linebreaks

import (
	"regexp"
)

// Map Map the Markdown input's line breaks into HTML elements
func Map(mdInput string) string {
	lineBreaksRegexp := regexp.MustCompile("(?m) {2,}$")

	return lineBreaksRegexp.ReplaceAllString(mdInput, "<br>")
}
