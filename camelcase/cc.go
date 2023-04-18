package camelcase

import (
	"unicode"
)

func camelcase(s string) int32 {
	result := 1

	for _, c := range s {
		if unicode.IsUpper(c) {
			result++
		}
	}

	return int32(result)
}
