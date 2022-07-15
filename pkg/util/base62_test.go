package util

import (
	"testing"
)

type base62EncodeTest struct {
	arg      int64
	expected string
	desc     string
}

func TestBase62Encode(t *testing.T) {
	base62 := Base62()
	testCases := []base62EncodeTest{
		{0, "0", "base62 format: [0-9][a-z][A-z]"},
		{1, "1", "returns 1 not 0"},
		{10, "a", "base62 format: [0-9][a-z][A-z]"},
		{36, "A", "base62 format: [0-9][a-z][A-z]"},
		{3425566475, "3JPkkb", "valid encoding"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if output, err := base62.Encode(tC.arg); output != tC.expected || err != nil {
				t.Fatalf("Expected %s, got %s instead.", tC.expected, output)
			}
		})
	}

	t.Run("when you enter negative value you should get an error", func(t *testing.T) {
		_, err := base62.Encode(-1)

		if err == nil {
			t.Fatal("Expected error")
		}
	})
}

func TestBase62Decode(t *testing.T) {

}
