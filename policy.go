package awsign

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"time"
)

var encode = base64.StdEncoding.EncodeToString

// helper types for these 2 structures
type (
	M map[string]string
	A []string
)

type Policy struct {
	Expiration time.Time     `json:"expiration"`
	Conditions []interface{} `json:"conditions"`
}

func NewPolicy(exp time.Time) *Policy {
	return &Policy{
		Expiration: exp,
	}
}

func (p *Policy) Add(v interface{}) {
	p.Conditions = append(p.Conditions, v)
}

func (p *Policy) Base64() (string, error) {
	b := bytes.NewBuffer(nil)

	err := json.NewEncoder(b).Encode(p)
	if err != nil {
		return "", err
	}

	return encode(b.Bytes()), nil
}

// String calls Base64, but suppresses any errors
func (p *Policy) String() string {
	str, _ := p.Base64()

	return str
}
