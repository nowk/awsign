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

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
}

type Options map[string]string

func (o Options) Set(k, v string) {
	o[k] = v
}

func (o Options) Get(k string) string {
	return o[k]
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
