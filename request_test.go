package awsign

import (
	"net/http"
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
		giv http.Header
		exp e
	}{
		{
			http.Header{}, e{``, ``},
		},
		{
			http.Header{
				"Host":         {"iam.amazonaws.com"},
				"Content-Type": {"application/x-www-form-urlencoded; charset=utf-8"},
			},
			e{
				"content-type;host",
				`content-type:application/x-www-form-urlencoded; charset=utf-8
host:iam.amazonaws.com
`},
		},
		{
			http.Header{
				"Host": {"    iam.amazonaws.com    "},
			},
			e{"host", "host:iam.amazonaws.com\n"},
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

func TestCanonicalRequest(t *testing.T) {
	var exp = `GET
/
Action=ListUsers&Version=2010-05-08
content-type:application/x-www-form-urlencoded; charset=utf-8
host:iam.amazonaws.com
x-amz-date:20150830T123600Z

content-type;host;x-amz-date
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`

	got := CanonicalRequest(testRequest)
	if exp != got {
		t.Errorf("expected %s, got %s", exp, got)
	}
}

func TestHashCanonicalRequest(t *testing.T) {
	var exp = "f536975d06c0309214f805bb90ccff089219ecd68b2577efef23edd43b7e1a59"

	got := HashCanonicalRequest(testRequest)
	if exp != got {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
