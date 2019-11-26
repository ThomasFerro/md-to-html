package emphasis_test

import (
	"testing"

	"github.com/ThomasFerro/md-to-html/emphasis"
)

func TestBoldTextAsterisks(t *testing.T) {
	textWithEmphasis := emphasis.Map("I just love **bold text**.")

	if textWithEmphasis != "I just love <strong>bold text</strong>." {
		t.Errorf("The bold text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}

func TestBoldTextUnderscores(t *testing.T) {
	textWithEmphasis := emphasis.Map("I just love __bold text__.")

	if textWithEmphasis != "I just love <strong>bold text</strong>." {
		t.Errorf("The bold text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}

func TestBoldTextMiddleOfAWord(t *testing.T) {
	textWithEmphasis := emphasis.Map("Love**is**bold")

	if textWithEmphasis != "Love<strong>is</strong>bold" {
		t.Errorf("The bold text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}

func TestItalicTextAsterisk(t *testing.T) {
	textWithEmphasis := emphasis.Map("Italicized text is the *cat's meow*.")

	if textWithEmphasis != "Italicized text is the <em>cat's meow</em>." {
		t.Errorf("The italic text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}

func TestItalicTextUnderscore(t *testing.T) {
	textWithEmphasis := emphasis.Map("Italicized text is the _cat's meow_.")

	if textWithEmphasis != "Italicized text is the <em>cat's meow</em>." {
		t.Errorf("The italic text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}

func TestItalicTextMiddleOfAWord(t *testing.T) {
	textWithEmphasis := emphasis.Map("A*cat*meow")

	if textWithEmphasis != "A<em>cat</em>meow" {
		t.Errorf("The italic text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}

func TestBoldAndItalicTextAsterisks(t *testing.T) {
	textWithEmphasis := emphasis.Map("This text is ***really important***.")

	if textWithEmphasis != "This text is <strong><em>really important</em></strong>." {
		t.Errorf("The bold and italic text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}

func TestBoldAndItalicTextUnderscores(t *testing.T) {
	textWithEmphasis := emphasis.Map("This text is ___really important___.")

	if textWithEmphasis != "This text is <strong><em>really important</em></strong>." {
		t.Errorf("The bold and italic text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}

func TestBoldAndItalicTextAsteriskAndUnderscores(t *testing.T) {
	textWithEmphasis := emphasis.Map("This text is __*really important*__.")

	if textWithEmphasis != "This text is <strong><em>really important</em></strong>." {
		t.Errorf("The bold and italic text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}

func TestBoldAndItalicTextAsterisksAndUnderscore(t *testing.T) {
	textWithEmphasis := emphasis.Map("This text is **_really important_**.")

	if textWithEmphasis != "This text is <strong><em>really important</em></strong>." {
		t.Errorf("The bold and italic text was not converted to HTML as intended, got \"%v\" instead", textWithEmphasis)
	}
}
