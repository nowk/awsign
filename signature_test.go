package awsign

import (
	"net/http"
	"testing"
)

const tSecret = "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY"

func TestSignature(t *testing.T) {
	req, err := http.NewRequest("GET", "https://iam.amazonaws.com/?Action=ListUsers&Version=2010-05-08", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("X-Amz-Date", "20150830T123600Z")
	req.Header.Add("Host", "iam.amazonaws.com")

	var exp = "5d672d79c15b13162d9279b0855cfba6789a8edb4c82c400e06b5924a6f2b5d7"

	confFunc := func(c Conf) {
		c.Set("service", "iam")
	}

	got := Signature(req, tSecret, confFunc)
	if exp != got {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
