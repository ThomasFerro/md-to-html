package horizontalrules_test

import (
	"testing"

	"github.com/ThomasFerro/md-to-html/horizontalrules"
)

func TestNoHorizontalRuleWhenLessThanThreeAsterisks(t *testing.T) {
	horizontalRule := horizontalrules.Map("**")

	if horizontalRule != "**" {
		t.Errorf("No horizontal rule should have be created, got \"%v\" instead", horizontalRule)
	}
}

func TestHorizontalRuleFromAsterisks(t *testing.T) {
	horizontalRule := horizontalrules.Map("***")

	if horizontalRule != "<hr>" {
		t.Errorf("The horizontal rule was not converted to HTML as intended, got \"%v\" instead", horizontalRule)
	}
}

func TestHorizontalRuleFromDashes(t *testing.T) {
	horizontalRule := horizontalrules.Map("-------")

	if horizontalRule != "<hr>" {
		t.Errorf("The horizontal rule was not converted to HTML as intended, got \"%v\" instead", horizontalRule)
	}
}

func TestHorizontalRuleFromUnderscores(t *testing.T) {
	horizontalRule := horizontalrules.Map("___________________")

	if horizontalRule != "<hr>" {
		t.Errorf("The horizontal rule was not converted to HTML as intended, got \"%v\" instead", horizontalRule)
	}
}

func TestMultipleHorizontalRules(t *testing.T) {
	horizontalRule := horizontalrules.Map(`___________________
---
***`)

	if horizontalRule != `<hr>
<hr>
<hr>` {
		t.Errorf("The horizontal rule was not converted to HTML as intended, got \"%v\" instead", horizontalRule)
	}
}
