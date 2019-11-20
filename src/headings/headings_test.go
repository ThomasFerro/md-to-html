package headings_test

import (
	"testing"

	"github.com/ThomasFerro/md-to-html/headings"
)

func TestHeadingLevelOne(t *testing.T) {
	header := headings.Map("# Heading level 1")

	if header != "<h1>Heading level 1</h1>" {
		t.Errorf("The heading was not converted to HTML as intended, got \"%v\" instead", header)
	}
}

func TestHeadingLevelTwo(t *testing.T) {
	header := headings.Map("## Heading level 2")

	if header != "<h2>Heading level 2</h2>" {
		t.Errorf("The heading was not converted to HTML as intended, got \"%v\" instead", header)
	}
}

func TestHeadingLevelThree(t *testing.T) {
	header := headings.Map("### Heading level 3")

	if header != "<h3>Heading level 3</h3>" {
		t.Errorf("The heading was not converted to HTML as intended, got \"%v\" instead", header)
	}
}

func TestHeadingLevelFour(t *testing.T) {
	header := headings.Map("#### Heading level 4")

	if header != "<h4>Heading level 4</h4>" {
		t.Errorf("The heading was not converted to HTML as intended, got \"%v\" instead", header)
	}
}

func TestHeadingLevelFive(t *testing.T) {
	header := headings.Map("##### Heading level 5")

	if header != "<h5>Heading level 5</h5>" {
		t.Errorf("The heading was not converted to HTML as intended, got \"%v\" instead", header)
	}
}

func TestHeadingLevelSix(t *testing.T) {
	header := headings.Map("###### Heading level 6")

	if header != "<h6>Heading level 6</h6>" {
		t.Errorf("The heading was not converted to HTML as intended, got \"%v\" instead", header)
	}
}

func TestNoHeadingAboveLevelSix(t *testing.T) {
	header := headings.Map("####### Heading level 7")

	if header != "####### Heading level 7" {
		t.Errorf("The heading was converted to HTML for a level above six, got \"%v\" instead", header)
	}
}

func TestMultipleHeadings(t *testing.T) {
	header := headings.Map(`# Heading level 1

Text level 1

## Heading level 2

Text level 2

## Another heading level 2

Another text level 2`)

	if header != `<h1>Heading level 1</h1>

Text level 1

<h2>Heading level 2</h2>

Text level 2

<h2>Another heading level 2</h2>

Another text level 2` {
		t.Errorf("Multiple headings are not converted, got \"%v\" instead", header)
	}
}

func TestHeadingLevelOneAlternateSyntax(t *testing.T) {
	header := headings.Map(`Heading level 1
===============`)

	if header != "<h1>Heading level 1</h1>" {
		t.Errorf("Alternate syntax for heading one is not converted, got \"%v\" instead", header)
	}
}

func TestHeadingLevelOneAlternateSyntaxIgnoredWhenNoTextIsPresent(t *testing.T) {
	header := headings.Map(`Not a heading level 1

===============`)

	if header != `Not a heading level 1

===============` {
		t.Errorf("Alternate syntax for heading one is not converted, got \"%v\" instead", header)
	}
}

func TestHeadingLevelTwoAlternateSyntax(t *testing.T) {
	header := headings.Map(`Heading level 2
---------------`)

	if header != "<h2>Heading level 2</h2>" {
		t.Errorf("Alternate syntax for heading two is not converted, got \"%v\" instead", header)
	}
}

func TestHeadingLevelTwoAlternateSyntaxIgnoredWhenNoTextIsPresent(t *testing.T) {
	header := headings.Map(`Not a heading level 2

---------------`)

	if header != `Not a heading level 2

---------------` {
		t.Errorf("Alternate syntax for heading two is not converted, got \"%v\" instead", header)
	}
}
