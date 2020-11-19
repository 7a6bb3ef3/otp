package otp

type OneTimePassword interface {
	Sum(k []byte) string
}