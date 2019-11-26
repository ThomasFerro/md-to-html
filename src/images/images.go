package images

import "regexp"

func replaceAllImages(input string) string {
	return regexp.MustCompile(`(?mU)!\[(.*)\]\((.*)\)`).ReplaceAllString(input, `<img src="$2" alt="$1">`)
}

func replaceAllImagesWithTitle(input string) string {
	return regexp.MustCompile(`(?mU)!\[(.*)\]\((.*) "(.*)"\)`).ReplaceAllString(input, `<img src="$2" alt="$1" title="$3">`)
}

// Map Map the Markdown input's images into HTML elements
func Map(mdInput string) string {
	output := replaceAllImagesWithTitle(mdInput)
	return replaceAllImages(output)
}
