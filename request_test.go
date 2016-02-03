package awsign

import (
	"testing"
)

func TestCanonicalURI(t *testing.T) {
	var cases = []struct {
		giv, exp string
	}{
		{"/", "/"},
		{"", "/"},
		{"/documents and settings", `/documents%20and%20settings`},
	}

	for _, v := range cases {
		var exp = v.exp

		got := CanonicalURI(v.giv)
		if exp != got {
			t.Errorf("expected %s, got %s", exp, got)
		}
	}
}
