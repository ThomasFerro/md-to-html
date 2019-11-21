package links

import (
	"regexp"
)

func replaceAll(input, pattern, substitution string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(input, substitution)
}

// Map Map the Markdown input's links into HTML elements
func Map(mdInput string) string {
	output := replaceAll(mdInput, `(?m)\[(.*)\]\((.*) "(.*)"\)`, `<a href="$2" title="$3" target="_blank">$1</a>`)

	return replaceAll(output, `(?m)\[(.*)\]\((.*)\)`, `<a href="$2" target="_blank">$1</a>`)
}
