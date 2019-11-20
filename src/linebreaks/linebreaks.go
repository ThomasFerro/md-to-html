package linebreaks

import (
	"regexp"
)

// Map Map the Markdown input's line breaks into HTML elements
func Map(mdInput string) string {
	lineBreaksRegexp := regexp.MustCompile("(?m)^(.+)  (\\n)")

	return lineBreaksRegexp.ReplaceAllString(mdInput, "$1<br>$2")
}
