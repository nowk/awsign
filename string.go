package awsign

import (
	"path"
)

// http://docs.aws.amazon.com/general/latest/gr/sigv4-create-string-to-sign.html

func StringToSign(algo, date, scope, hash string) string {
	return join([]string{
		algo,
		date,
		scope,
		hash,
	}, "\n")
}

func CredentialScope(s ...string) string {
	return path.Join(append(s, "aws4_request")...)
}
