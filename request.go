package awsign

import (
	"net/url"
	"sort"
	"strings"
)

var (
	lower = strings.ToLower
	join  = strings.Join
	trim  = strings.TrimSpace
)

// http://docs.aws.amazon.com/general/latest/gr/sigv4-create-canonical-request.html

func CanonicalURI(uri string) string {
	if uri == "" || uri == "/" {
		return "/"
	}
	u := url.URL{Path: uri}

	return u.String()
}

// CanonicalHeaders returns both the signed headers and the canonical header
// entires
func CanonicalHeaders(h url.Values) (string, string) {
	lenh := len(h)
	if lenh == 0 {
		return "", ""
	}

	var (
		so = make([]string, 0, lenh)
		en = make([]string, 0, lenh)
	)

	// get headers and sort
	for k, _ := range h {
		so = append(so, k)
	}
	sort.Strings(so)

	for _, v := range so {
		val := h.Get(v)
		en = append(en, lower(v)+":"+trim(val))
	}

	var (
		sh = lower(join(so, ";")) + "\n"
		he = join(en, "\n") + "\n"
	)
	return sh, he
}
