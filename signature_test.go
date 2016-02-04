package awsign

import (
	"testing"
)

const tSecret = "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY"

func TestSignature(t *testing.T) {
	var exp = "5d672d79c15b13162d9279b0855cfba6789a8edb4c82c400e06b5924a6f2b5d7"

	var (
		c = CredentialScope("20150830", "us-east-1", "iam")
		h = HashCanonicalRequest(testRequest)
		s = StringToSign(defaultAlgo, "20150830T123600Z", c, h)
	)

	got := Signature(s, tSecret, "20150830", "us-east-1", "iam")
	if exp != got {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
