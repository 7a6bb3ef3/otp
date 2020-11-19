package hotp

import (
	"testing"
)

// RFC4226
var cases = map[uint64]string{
	0: "755224",
	1: "287082",
	2: "359152",
	3: "969429",
	4: "338314",
	5: "254676",
}

var secret = []byte("12345678901234567890")

func TestNew(t *testing.T) {
	for k, v := range cases {
		if h := New(secret, k); h != v {
			t.Log(k, v, h)
			t.Fail()
		} else {
			t.Log("Pass:", k, v, h)
		}
	}
}

func TestPow10(t *testing.T) {
	cases := make(map[uint]uint64)
	cases[0] = 1
	cases[1] = 10
	cases[4] = 10000
	for k, v := range cases {
		if n := pow10(k); n != v {
			t.Log(k, n, v)
			t.Fail()
		}
	}
}
