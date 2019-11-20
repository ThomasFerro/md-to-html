package horizontalrules

import (
	"regexp"
)

// Map Map the Markdown input's horizontal rules into HTML elements
func Map(mdInput string) string {
	horizontalRuleRegexp := regexp.MustCompile("(?m)^[*-_]{3,}$")
	return horizontalRuleRegexp.ReplaceAllString(mdInput, "<hr>")
}
