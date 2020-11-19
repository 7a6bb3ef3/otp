package hotp

import "hash"

type Option func(h *HOTP)

func Digit(n uint) Option {
	return func(h *HOTP) {
		h.Digit = n
	}
}

func Hash(hs func()hash.Hash) Option {
	return func(h *HOTP) {
		h.Hash = hs
	}
}
