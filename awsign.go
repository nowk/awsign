package awsign

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net/http"
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

func Sign(req *http.Request, c Credentials) Options {
	var (
		u         = req.URL
		h         = req.Header
		date      = h.Get("X-Amz-Date")
		dateshort = date[:8]
	)
	var (
		hash  = HashCanonicalRequest(req)
		scope = CredentialScope(dateshort, "us-east-1", "iam")
		str   = StringToSign(defaultAlgo, date, scope, hash)
		sig   = Signature(str, c.SecretAccessKey, dateshort, "us-east-1", "iam")
	)

	return Options{
		"key":        u.Path,
		"algorithm":  defaultAlgo,
		"credential": path.Join(c.AccessKeyID, scope),
		"signature":  sig,
		"date":       date,
	}
}
