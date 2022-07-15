package util

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

const (
	base         = 62
	characterSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base62Regex  = "^[0-9A-Za-z]+$"
)

type base62 struct {
	characterSet string
	base         int64
}

func (b *base62) Encode(num int64) (string, error) {
	if num == 0 {
		return "0", nil
	} else if num < 0 {
		return "", fmt.Errorf("number: %d must be greater or equal to zero", num)
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

	return result, nil
}

func (b *base62) Decode(str string) (int64, error) {
	if !isValidBase62(str) {
		return -1, fmt.Errorf("base62 doesn't match the pattern %s", base62Regex)
	}

	var result int64
	var exp int

	for i := len(str) - 1; i >= 0; i-- {
		if idx := strings.IndexByte(b.characterSet, str[i]); idx >= 0 {
			x := float64(b.base)
			y := float64(exp)
			temp := math.Pow(x, y)
			result += (int64(idx) * int64(temp))
			exp += 1
		}
	}

	return result, nil
}

func isValidBase62(s string) bool {
	match, _ := regexp.MatchString(base62Regex, s)
	return match
}

func Base62() *base62 {
	return &base62{
		characterSet: characterSet,
		base:         base,
	}
}
