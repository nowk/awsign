package awsign

import (
	"net/http"
	"testing"
)

func TestStringToSign(t *testing.T) {
	req, err := http.NewRequest("GET", "https://iam.amazonaws.com/?Action=ListUsers&Version=2010-05-08", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("X-Amz-Date", "20150830T123600Z")
	req.Header.Add("Host", "iam.amazonaws.com")

	var exp = `AWS4-HMAC-SHA256
20150830T123600Z
20150830/us-east-1/iam/aws4_request
f536975d06c0309214f805bb90ccff089219ecd68b2577efef23edd43b7e1a59`

	confFunc := func(c Conf) {
		c.Set("service", "iam")
	}

	got := StringToSign(req, confFunc)
	if exp != got {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
