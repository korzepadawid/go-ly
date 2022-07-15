package util

import "strings"

const (
	base         = 62
	characterSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type base62 struct {
	characterSet string
	base         int64
}

func (b *base62) Encode(num int64) string {
	var sb strings.Builder
	for num >= 0 {
		idx := int(num % base)
		char := string(b.characterSet[idx])
		sb.WriteString(char)
		num /= b.base
	}
	return sb.String()
}

func (b *base62) Decode(s string) (int64, error) {
	return int64(4), nil
}

func Base62() *base62 {
	return &base62{
		characterSet: characterSet,
		base:         base,
	}
}
