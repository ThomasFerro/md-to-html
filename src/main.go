package main

import (
	"github.com/ThomasFerro/md-to-html/headings"
)

// MarkdownToHTML Process the Markdown to create a HTML string
func MarkdownToHTML(markdown string) string {
	return headings.Map(markdown)
}
