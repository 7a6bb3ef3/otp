package totp

import "hash"

type Option func(totp *TOTP)

func T0(t0 int64) Option {
	return func(totp *TOTP) {
		totp.T0 = t0
	}
}

func T(t int64) Option {
	return func(totp *TOTP) {
		totp.T = t
	}
}

func X(x uint) Option {
	return func(totp *TOTP) {
		totp.X = x
	}
}

func Digit(d uint) Option {
	return func(totp *TOTP) {
		totp.Digit = d
	}
}

func Hash(h func() hash.Hash) Option {
	return func(totp *TOTP) {
		totp.Hash = h
	}
}

