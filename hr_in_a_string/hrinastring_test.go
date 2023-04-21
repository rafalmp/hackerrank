package hrinastring

import "testing"

func TestHackerrankInString(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"haacckkerrannkk", "YES"},
		{"haacckkerannk", "NO"},
		{"hccaakkerrannkk", "NO"},
		{"hereiamstackerrank", "YES"},
		{"hackerworld", "NO"},
		{"rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt", "NO"},
		{"hackerrank", "YES"},
		{"aaahackerrank", "YES"},
		{"hackerrankabc", "YES"},
	}

	for _, tst := range tests {
		t.Run(tst.input, func(t *testing.T) {
			got := hackerrankInString(tst.input)
			if got != tst.want {
				t.Errorf("got %s, want %s", got, tst.want)
			}
		})
	}
}
