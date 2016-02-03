package awsign

import (
	"net/url"
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

func TestCanonicalHeaders(t *testing.T) {
	type e struct {
		headers, entries string
	}

	var cases = []struct {
		giv url.Values
		exp e
	}{
		{
			url.Values{}, e{``, ``},
		},
		{
			url.Values{
				"Host":         {"iam.amazonaws.com"},
				"Content-Type": {"application/x-www-form-urlencoded; charset=utf-8"},
			},
			e{
				"content-type;host\n",
				`content-type:application/x-www-form-urlencoded; charset=utf-8
host:iam.amazonaws.com
`},
		},
		{
			url.Values{
				"Host": {"    iam.amazonaws.com    "},
			},
			e{"host\n", "host:iam.amazonaws.com\n"},
		},
	}

	for _, v := range cases {
		hd, en := CanonicalHeaders(v.giv)

		{
			var exp = v.exp.headers

			if got := hd; exp != got {
				t.Errorf("expected %s, got %s", exp, got)
			}
		}

		{
			var exp = v.exp.entries

			if got := en; exp != got {
				t.Errorf("expected %s, got %s", exp, got)
			}
		}
	}
}
