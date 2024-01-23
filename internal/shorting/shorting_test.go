package shorting

import (
	"strings"
	"testing"
)

func TestGenerateShortLinkLength(t *testing.T) {
	shortLink := GenerateShortLink()
	shortLink2 := GenerateShortLink()
	if len(shortLink) != Size || len(shortLink2) != Size {
		t.Errorf("Invalid short link length")
	}
	if shortLink == shortLink2 {
		t.Errorf("Short links are not generated randomly")
	}
	if len(strings.Trim(shortLink, Alphabet)) != 0 {
		t.Errorf("Short link contains invalid characters")
	}
}
