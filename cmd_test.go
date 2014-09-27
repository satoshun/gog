package main

import (
	"testing"
)

func TestOriginUrl(t *testing.T) {
	expectURL := "https://github.com/satoshun/gog"
	path := OriginUrl(".")
	if path != expectURL {
		t.Errorf("%s is not %s", path, expectURL)
	}
}
