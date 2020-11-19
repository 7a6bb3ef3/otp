package totp

import (
	"testing"
)

var cas map[int64]string
var key string

func init() {
	key = "12345678901234567890"
	cas = make(map[int64]string)
	cas[1111111109] = "07081804"
	cas[1111111111] = "14050471"
	cas[2000000000] = "69279037"
}

func TestNew(t *testing.T) {
	sec := []byte(key)
	for k, v := range cas {
		if r := New(sec, Digit(8), T(k)); r != v {
			t.Log(k, v, r)
			t.Fail()
		} else {
			t.Log("Pass:", k, v, r)
		}
	}
}

// 197874
// 424890
