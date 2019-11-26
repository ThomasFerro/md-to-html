package images_test

import (
	"testing"

	"github.com/ThomasFerro/md-to-html/images"
)

func TestImage(t *testing.T) {
	image := images.Map(`![Philadelphia's Magic Gardens. This place was so cool!](/assets/images/philly-magic-gardens.jpg)`)

	if image != `<img src="/assets/images/philly-magic-gardens.jpg" alt="Philadelphia's Magic Gardens. This place was so cool!">` {
		t.Errorf("The image was not converted to HTML as intended, got \"%v\" instead", image)
	}
}

func TestImageWithTitle(t *testing.T) {
	image := images.Map(`![Philadelphia's Magic Gardens. This place was so cool!](/assets/images/philly-magic-gardens.jpg "Philadelphia's Magic Gardens")`)

	if image != `<img src="/assets/images/philly-magic-gardens.jpg" alt="Philadelphia's Magic Gardens. This place was so cool!" title="Philadelphia's Magic Gardens">` {
		t.Errorf("The image was not converted to HTML as intended, got \"%v\" instead", image)
	}
}
