package main

import (
	"github.com/ThomasFerro/md-to-html/emphasis"
	"github.com/ThomasFerro/md-to-html/headings"
	"github.com/ThomasFerro/md-to-html/horizontalrules"
	"github.com/ThomasFerro/md-to-html/images"
	"github.com/ThomasFerro/md-to-html/linebreaks"
	"github.com/ThomasFerro/md-to-html/links"
)

// MarkdownToHTML Process the Markdown to create a HTML string
func MarkdownToHTML(markdown string) string {
	output := headings.Map(markdown)
	output = linebreaks.Map(output)
	output = horizontalrules.Map(output)
	output = images.Map(output)
	output = links.Map(output)
	output = emphasis.Map(output)
	return output
}
