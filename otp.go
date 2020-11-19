package otp

import (
	"encoding/base32"
	"fmt"
)

type OneTimePassword interface {
	Sum(k []byte) string
}

type Type string

const (
	TypeHOTP Type = "hotp"
	TypeTOTP Type = "totp"
)

type Algorithm string

const (
	SHA1   Algorithm = "SHA1"
	SHA256 Algorithm = "SHA256"
	SHA512 Algorithm = "SHA512"
)

// URI otpauth://TYPE/LABEL?PARAMETERS
// https://github.com/google/google-authenticator/wiki/Key-Uri-Format
// and test here https://rootprojects.org/authenticator/
type URI struct {
	Type      Type
	Issuer    string
	Account   string
	Secret    string
	Algorithm Algorithm
	Digit     uint
	Period    uint
	Counter   uint
}

// Default will set secret/algorithm/digit/period/counter to default value.
func Default() URI {
	return URI{
		Type:      "",
		Issuer:    "",
		Account:   "",
		Secret:    "",
		Algorithm: SHA1,
		Digit:     6,
		Period:    30,
		Counter:   0,
	}
}

// String will encode Secret to base32 format.
func (u URI) String() string {
	return fmt.Sprintf("otpauth://%s/%s:%s?"+
		"secret=%s&"+
		"algorithm=%s&"+
		"digit=%d&"+
		"period=%d&"+
		"counter=%d", u.Type, u.Issuer, u.Account, u.Base32Secret(), u.Algorithm, u.Digit, u.Period, u.Counter)
}

func (u URI) Base32Secret() string {
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString([]byte(u.Secret))
}
