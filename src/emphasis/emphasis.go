package emphasis

import "regexp"

func replaceAll(input, pattern, substitution string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(input, substitution)
}

func replaceAllBoldAndItalic(input string) string {
	substitution := "<strong><em>$1</em></strong>"
	output := replaceAll(input, `(?mU)\*\*\*(.*)\*\*\*`, substitution)
	return replaceAll(output, `(?mU)___(.*)___`, substitution)
}

func replaceAllBold(input string) string {
	substitution := "<strong>$1</strong>"
	output := replaceAll(input, `(?mU)__(.*)__`, substitution)
	return replaceAll(output, `(?mU)\*\*(.*)\*\*`, substitution)
}

func replaceAllItalic(input string) string {
	substitution := "<em>$1</em>"
	output := replaceAll(input, `(?mU)\*(.*)\*`, substitution)
	return replaceAll(output, `(?mU)_(.*)_`, substitution)
}

// Map Map the Markdown input's emphasis into HTML elements
func Map(mdInput string) string {
	output := replaceAllBoldAndItalic(mdInput)
	output = replaceAllBold(output)
	return replaceAllItalic(output)
}
