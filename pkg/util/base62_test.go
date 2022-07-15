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

	t.Run("if you enter negative value, you should get an error", func(t *testing.T) {
		_, err := base62.Encode(-1)
		if err == nil {
			t.Fatal("Expected error")
		}
	})
}

type base62DecodeError struct {
	arg  string
	desc string
}

type base62Decode struct {
	arg      string
	expected int64
	desc     string
}

func TestBase62Decode(t *testing.T) {
	base62 := Base62()

	errorTestCases := []base62DecodeError{
		{"12hę", "if you use diacritical sign, you should get an error"},
		{"helL-0", "if you use dash sign, you should get an error"},
	}

	for _, tC := range errorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			if _, err := base62.Decode(tC.arg); err == nil {
				t.Fatal("Expected error")
			}
		})
	}

	happyPathTestCases := []base62Decode{
		{"3JPkkb", 3425566475, "if you enter valid base62, you should get decimal result"},
		{"z2xJK", 517778104, "if you enter valid base62, you should get decimal result"},
		{"0", 0, "if you enter valid base62, you should get decimal result"},
		{"a", 10, "if you enter valid base62, you should get decimal result"},
		{"a", 10, "if you enter valid base62, you should get decimal result"},
		{"A", 36, "if you enter valid base62, you should get decimal result"},
		{"Z", 61, "if you enter valid base62, you should get decimal result"},
	}

	for _, tC := range happyPathTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			if output, err := base62.Decode(tC.arg); output != tC.expected || err != nil {
				t.Fatalf("Expected %d, got %d instead", tC.expected, output)
			}
		})
	}
}
