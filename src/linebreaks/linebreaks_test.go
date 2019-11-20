package linebreaks_test

import (
	"testing"

	"github.com/ThomasFerro/md-to-html/linebreaks"
)

func TestLineBreaks(t *testing.T) {
	lineBreak := linebreaks.Map(`This is the first line.  
And this is the second line.`)

	if lineBreak != `This is the first line.<br>
And this is the second line.` {
		t.Errorf("The line break was not converted to HTML as intended, got \"%v\" instead", lineBreak)
	}
}

func TestLineBreaksMoreThanTwoSpaces(t *testing.T) {
	lineBreak := linebreaks.Map(`This is the first line.   
And this is the second line.`)

	if lineBreak != `This is the first line.<br>
And this is the second line.` {
		t.Errorf("The line break was not converted to HTML as intended, got \"%v\" instead", lineBreak)
	}
}
