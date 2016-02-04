package awsign

import (
	"encoding/hex"
	"net/http"
)

var toString = hex.EncodeToString

// http://docs.aws.amazon.com/general/latest/gr/sigv4-calculate-signature.html

func SigningKey(secret, date string, c Conf) []byte {
	var (
		region  = []byte(c.Get("region"))
		service = []byte(c.Get("service"))

		kdate    = hmacSha256([]byte("AWS4"+secret), []byte(date))
		kregion  = hmacSha256(kdate, region)
		kservice = hmacSha256(kregion, service)
	)

	return hmacSha256(kservice, []byte("aws4_request"))
}

func Signature(req *http.Request, secret string, opts ...func(Conf)) string {
	var (
		h         = req.Header
		date      = h.Get("X-Amz-Date")
		str, conf = StringToSign(req, opts...)

		key = SigningKey(secret, date[:8], conf)
	)

	return toString(hmacSha256(key, []byte(str)))
}
