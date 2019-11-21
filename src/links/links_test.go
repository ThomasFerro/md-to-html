package links_test

import (
	"testing"

	"github.com/ThomasFerro/md-to-html/links"
)

func TestLink(t *testing.T) {
	link := links.Map("My favorite search engine is [Duck Duck Go](https://duckduckgo.com).")

	if link != `My favorite search engine is <a href="https://duckduckgo.com" target="_blank">Duck Duck Go</a>.` {
		t.Errorf("The link was not converted to HTML as intended, got \"%v\" instead", link)
	}
}

func TestLinkWithTitle(t *testing.T) {
	link := links.Map(`My favorite search engine is [Duck Duck Go](https://duckduckgo.com "The best search engine for privacy").`)

	if link != `My favorite search engine is <a href="https://duckduckgo.com" title="The best search engine for privacy" target="_blank">Duck Duck Go</a>.` {
		t.Errorf("The link was not converted to HTML as intended, got \"%v\" instead", link)
	}
}
