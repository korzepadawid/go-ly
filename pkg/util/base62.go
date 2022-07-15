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
	if num == 0 {
		return "0"
	}

	var result string

	for num > 0 {
		idx := int(num % b.base)
		num /= b.base
		sb := strings.Builder{}
		sb.WriteString(string(characterSet[idx]))
		sb.WriteString(result)
		result = sb.String()
	}

	return result
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