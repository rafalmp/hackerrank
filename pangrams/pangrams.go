package pangrams

import "unicode"

func pangrams(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	chars := make(map[rune]struct{})

	for _, c := range alphabet {
		chars[c] = struct{}{}
	}

	for _, c := range s {
		delete(chars, unicode.ToLower(c))
		if len(chars) == 0 {
			return "pangram"
		}
	}
	return "not pangram"
}
