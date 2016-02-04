package awsign

import (
	"net/http"
	"reflect"
	"testing"
)

// testRequest is modelled after this request
//
//   GET https://iam.amazonaws.com/?Action=ListUsers&Version=2010-05-08 HTTP/1.1
//   Host: iam.amazonaws.com
//   Content-Type: application/x-www-form-urlencoded; charset=utf-8
//   X-Amz-Date: 20150830T123600Z
//
var testRequest = func() *http.Request {
	req, err := http.NewRequest("GET", "https://iam.amazonaws.com/?Action=ListUsers&Version=2010-05-08", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("X-Amz-Date", "20150830T123600Z")
	req.Header.Add("Host", "iam.amazonaws.com")

	return req
}()

func TestSign(t *testing.T) {
	var exp = Options{
		"key":        "/",
		"algorithm":  defaultAlgo,
		"credential": "0123456789/20150830/us-east-1/iam/aws4_request",
		"signature":  "5d672d79c15b13162d9279b0855cfba6789a8edb4c82c400e06b5924a6f2b5d7",
		"date":       "20150830T123600Z",
	}

	c := Credentials{
		AccessKeyID:     "0123456789",
		SecretAccessKey: tSecret,
	}

	got := Sign(testRequest, c)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
