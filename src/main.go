package main

import (
	"github.com/ThomasFerro/md-to-html/headings"
	"github.com/ThomasFerro/md-to-html/linebreaks"
)

// MarkdownToHTML Process the Markdown to create a HTML string
func MarkdownToHTML(markdown string) string {
	output := headings.Map(markdown)
	output = linebreaks.Map(output)
	return output
}
