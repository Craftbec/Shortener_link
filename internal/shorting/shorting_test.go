package shorting

import (
	"strings"
	"testing"
)

func TestGenerateShortLinkLength(t *testing.T) {
	shortLink := GenerateShortLink()
	shortLink2 := GenerateShortLink()
	if len(shortLink) != Size || len(shortLink2) != Size {
		t.Error("Invalid short link length\n")
	}
	if shortLink == shortLink2 {
		t.Error("Short links are not generated randomly\n")
	}
	if len(strings.Trim(shortLink, Alphabet)) != 0 {
		t.Error("Short link contains invalid characters\n")
	}
}
