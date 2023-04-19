package caesar

import (
	"strings"
	"unicode"
)

func caesarCipher(s string, k int32) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	alphabetLength := len(alphabet)
	var result string
	for _, c := range s {
		if i := strings.IndexRune(alphabet, unicode.ToLower(c)); i != -1 {
			rot := string(alphabet[(i+int(k))%alphabetLength])
			if unicode.IsUpper(c) {
				rot = strings.ToUpper(rot)
			}
			result += rot
			continue
		}
		result += string(c)
	}
	return result
}
