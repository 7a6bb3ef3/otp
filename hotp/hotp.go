package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"hash"
	"strconv"
)

// RFC4226
type HOTP struct {
	// RFC4226 Section 5.1. shared secret between client and server; each HOTP
	// generator has a different and unique secret K.
	K []byte
	// RFC4226 Section 5.1. 8-byte counter value, the moving factor.  This counter
	// MUST be synchronized between the HOTP generator (client)
	// and the HOTP validator (server).
	C uint64
	// RFC4226 Section5.1. Default 6. number of digits in an HOTP value; system parameter.
	Digit uint
	Hash  func() hash.Hash
}

func (h *HOTP) Sum(k []byte) string {
	// RFC4226 Section 5.3 Step 1
	bf := make([]byte, 8)
	binary.BigEndian.PutUint64(bf, h.C)
	hm := hmac.New(h.Hash, k)
	hm.Write(bf)
	hs := hm.Sum(nil)

	// Step 2
	offset := hs[19] & 0xf
	sbit := hs[offset : offset+4]
	sbit[0] = sbit[0] & 0x7f
	// Step 3
	for k := range bf {
		bf[k] = 0
	}
	bf[4] = sbit[0]
	bf[5] = sbit[1]
	bf[6] = sbit[2]
	bf[7] = sbit[3]
	u := binary.BigEndian.Uint64(bf)
	format := "%0" + strconv.Itoa(int(h.Digit)) + "d"
	return fmt.Sprintf(format, u%pow10(h.Digit))
}

func New(key []byte, c uint64, opts ...Option) string {
	h := &HOTP{
		K:     key,
		C:     c,
		Digit: 6,
		Hash:  sha1.New,
	}
	for _, f := range opts {
		f(h)
	}
	return h.Sum(key)
}

func pow10(n uint) uint64 {
	var r uint64 = 1
	for i := 0; i < int(n); i++ {
		r *= 10
	}
	return r
}
