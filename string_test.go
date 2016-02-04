package awsign

import (
	"net/http"
	"reflect"
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

	confFunc := func(c Conf) {
		c.Set("service", "iam")
	}
	s, c := StringToSign(req, confFunc)

	{
		var exp = `AWS4-HMAC-SHA256
20150830T123600Z
20150830/us-east-1/iam/aws4_request
f536975d06c0309214f805bb90ccff089219ecd68b2577efef23edd43b7e1a59`

		if got := s; exp != got {
			t.Errorf("expected %s, got %s", exp, got)
		}
	}

	{
		var exp = Conf{
			"algo":    defaultAlgo,
			"service": "iam",
			"region":  "us-east-1",
		}

		if got := c; !reflect.DeepEqual(exp, got) {
			t.Errorf("expected %s, got %s", exp, got)
		}
	}
}
