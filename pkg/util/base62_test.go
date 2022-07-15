package util

import (
	"fmt"
	"testing"
)

type base62EncodeTest struct {
	arg      int64
	expected string
	desc     string
}

func TestBase62Encode(t *testing.T) {
	testCases := []base62EncodeTest{
		{0, "0", "base62 format: [0-9][a-z][A-z]"},
		{1, "1", "returns 1 not 0"},
		{10, "a", "base62 format: [0-9][a-z][A-z]"},
		{36, "A", "base62 format: [0-9][a-z][A-z]"},
		{3425566475, "3JPkkb", "valid encoding"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			base62 := Base62()
			fmt.Println(base62.Encode(tC.arg))
		})
	}
}
