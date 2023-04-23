package pangrams

import "testing"

func TestPangrams(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"The quick brown fox jumps over the lazy dog", "pangram"},
		{"We promptly judged antique ivory buckles for the next prize", "pangram"},
		{"We promptly judged antique ivory buckles for the prize", "not pangram"},
		{"abcdefghijklmnopqrstuvwxyz", "pangram"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "pangram"},
	}

	for _, tst := range tests {
		t.Run(tst.input, func(t *testing.T) {
			got := pangrams(tst.input)
			if got != tst.want {
				t.Errorf("got %s, want %s", got, tst.want)
			}
		})
	}
}
