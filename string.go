package awsign

import (
	"net/http"
	"path"
)

func defaultConf() Conf {
	return Conf{
		"algo":    defaultAlgo,
		"service": "s3",
		"region":  "us-east-1",
	}
}

// http://docs.aws.amazon.com/general/latest/gr/sigv4-create-string-to-sign.html

func StringToSign(req *http.Request, opts ...func(Conf)) string {
	var (
		h    = req.Header
		date = h.Get("X-Amz-Date")
	)
	// TODO error if date is empty, bad format (or does not meet a basic length)

	conf := defaultConf()
	for _, v := range opts {
		v(conf)
	}
	var (
		algo = conf.Get("algo")

		// TODO service and region should be parsed from the endpoint url
		service = conf.Get("service")
		region  = conf.Get("region")
	)

	s := []string{
		algo,
		date,
		CredentialScope(date[:8], region, service),
		HashCanonicalRequest(req),
	}

	return join(s, "\n")
}

func CredentialScope(s ...string) string {
	return path.Join(append(s, "aws4_request")...)
}
