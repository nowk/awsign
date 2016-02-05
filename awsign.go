package awsign

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"path"
)

func hashSha256(b []byte) string {
	sha := sha256.New()
	sha.Write(b)
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func hmacSha256(key, content []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(content)
	return h.Sum(nil)
}

const defaultAlgo = "AWS4-HMAC-SHA256"

type Options map[string]string

func (o Options) Set(k, v string) {
	o[k] = v
}

func (o Options) Get(k string) string {
	return o[k]
}

type Aws map[string]string

func (a Aws) Set(k, v string) {
	a[k] = v
}

func (a Aws) Get(k string) string {
	return a[k]
}

var defaultAws = func() Aws {
	return Aws{
		"algorithm": defaultAlgo,
		"region":    "us-east-1",
		"service":   "s3",
	}
}

// func Sign(req *http.Request, c Credentials) Options {
// 	var (
// 		h         = req.Header
// 		date      = h.Get("X-Amz-Date")
// 		dateshort = date[:8]

// 		region  = "us-east-1"
// 		service = "s3"

// 		hash      = HashCanonicalRequest(req)
// 		scope     = CredentialScope(dateshort, region, service)
// 		strtosign = StringToSign(defaultAlgo, date, scope, hash)

// 		credential = path.Join(c.AccessKeyID, scope)
// 	)

// 	sig := Signature(strtosign, c.SecretAccessKey, dateshort, region, service)

// 	o := Options{
// 		"algorithm":  defaultAlgo,
// 		"credential": credential,
// 		"signature":  sig,
// 		"date":       date,
// 	}

// 	return o
// }

// SignPolicy is used to create signatures using the Poicy as the string to sign
// as documented for AWS S3 browser based uploading.
// http://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-UsingHTTPPOST.html
func SignPolicy(
	p *Policy, date string, c Aws, opts ...func(Aws)) (Options, error) {

	aws := defaultAws()
	for _, v := range opts {
		v(aws)
	}
	var (
		algo    = aws.Get("algorithm")
		region  = aws.Get("region")
		service = aws.Get("service")

		dateshort = date[:8]

		credential = path.Join(
			c.Get("access_key_id"), CredentialScope(dateshort, region, service))
	)

	strtosign, err := p.Base64()
	if err != nil {
		return nil, err
	}

	sig := Signature(strtosign, c.Get("secret_access_key"), dateshort, region, service)

	o := Options{
		"algorithm":  algo,
		"credential": credential,
		"policy":     strtosign,
		"signature":  sig,
		"date":       date,
	}

	return o, nil
}
