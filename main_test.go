package funpics

import (
	"testing"
)

func TestProviders(t *testing.T) {
	var url string
	url = GetFrom(0)
	if len(url) == 0 {
		t.Error("Expected url from provider2", url)
	}

	url = GetFrom(1)
	if len(url) == 0 {
		t.Error("Expected url from provider1 ", url)
	}
}
