package awsign

import (
	"testing"
)

func TestStringToSign(t *testing.T) {
	var (
		c = CredentialScope("20150830", "us-east-1", "iam")
		h = HashCanonicalRequest(testRequest)
	)

	s := StringToSign(defaultAlgo, "20150830T123600Z", c, h)

	var exp = `AWS4-HMAC-SHA256
20150830T123600Z
20150830/us-east-1/iam/aws4_request
f536975d06c0309214f805bb90ccff089219ecd68b2577efef23edd43b7e1a59`

	if got := s; exp != got {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
