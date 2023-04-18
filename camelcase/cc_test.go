package camelcase

import (
	"testing"
)

func TestCamelcase(t *testing.T) {
	var tests = []struct {
		input string
		want  int32
	}{
		{"saveChangesInTheEditor", 5},
		{"a", 1},
	}

	for _, tst := range tests {
		t.Run(tst.input, func(t *testing.T) {
			got := camelcase(tst.input)
			if got != tst.want {
				t.Errorf("got %d, want %d", got, tst.want)
			}
		})
	}
}
