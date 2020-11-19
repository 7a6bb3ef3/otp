package otp

import (
	"github.com/nynicg/otp/totp"
	"testing"
	"time"
)

func TestURI_String(t *testing.T) {
	uri := URI{
		Type:      TypeTOTP,
		Issuer:    "Cookie",
		Account:   "RRM",
		Secret:    "Hello RRM Aneki",
		Algorithm: SHA1,
		Digit:     6,
		Period:    30,
		Counter:   0,
	}
	t.Log(totp.New([]byte(uri.Secret)))
	tk := time.NewTicker(time.Second * time.Duration(uri.Period))
	defer tk.Stop()
	for range tk.C {
		code := totp.New([]byte(uri.Secret))
		t.Log(code)
	}
}
