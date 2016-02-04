package awsign

import (
	"encoding/hex"
)

var toString = hex.EncodeToString

// http://docs.aws.amazon.com/general/latest/gr/sigv4-calculate-signature.html

func SigningKey(secret, date, region, service string) []byte {
	var (
		kdate    = hmacSha256([]byte("AWS4"+secret), []byte(date))
		kregion  = hmacSha256(kdate, []byte(region))
		kservice = hmacSha256(kregion, []byte(service))
	)

	return hmacSha256(kservice, []byte("aws4_request"))
}

func Signature(s, secret, date, region, service string) string {
	return toString(
		hmacSha256(SigningKey(secret, date, region, service), []byte(s)))
}
