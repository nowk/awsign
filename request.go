package awsign

import (
	"net/url"
)

// http://docs.aws.amazon.com/general/latest/gr/sigv4-create-canonical-request.html

func CanonicalURI(uri string) string {
	if uri == "" || uri == "/" {
		return "/"
	}
	u := url.URL{Path: uri}

	return u.String()
}
