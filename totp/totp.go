package totp

import (
	"crypto/sha1"
	"github.com/nynicg/otp/hotp"
	"hash"
	"time"
)

// RFC6238 Basically, we define TOTP as TOTP = HOTP(K, T).
// TOTP implementations MAY use HMAC-SHA-1 ,HMAC-SHA-256 or HMAC-SHA-512 functions,
type TOTP struct {
	// In most cases ,T should be current unix time.
	T			int64
	// T0 is the Unix time to start counting time steps (default value is
	// 0, i.e., the Unix epoch) and is also a system parameter.
	T0			int64
	// represents the time step in seconds (default value X =
	// 30 seconds) and is a system parameter.
	X			uint
	// number of digits in an HOTP value; system parameter.
	Digit		uint
	// HMAC-SHA-1 ,HMAC-SHA-256 or HMAC-SHA-512
	Hash		func() hash.Hash
}

func New(key []byte ,opts ...Option) string{
	ttp := &TOTP{
		T0: 0,
		X:  30,
		Hash: sha1.New,
		Digit: 6,
		T:time.Now().Unix(),
	}
	for _ ,f := range opts{
		f(ttp)
	}
	c := (ttp.T - ttp.T0) / int64(ttp.X)
	return hotp.New(key ,uint64(c) , hotp.Hash(ttp.Hash) , hotp.Digit(ttp.Digit))
}