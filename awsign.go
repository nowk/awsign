package awsign

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
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

type Conf map[string]string

func (c Conf) Set(k, v string) {
	c[k] = v
}

func (c Conf) Get(k string) string {
	return c[k]
}
