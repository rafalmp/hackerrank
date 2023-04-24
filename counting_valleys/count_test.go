package countingvalleys

import (
	"testing"
)

func TestCountingValleys(t *testing.T) {
	var tests = []struct {
		steps int32
		path  string
		want  int32
	}{
		{8, "DDUUUUDD", 1},
		{8, "UDDDUDUU", 1},
		{10, "UDDUDUUDDU", 3},
	}

	for _, tst := range tests {
		t.Run(tst.path, func(t *testing.T) {
			got := countingValleys(tst.steps, tst.path)
			if got != tst.want {
				t.Errorf("got %d, want %d", got, tst.want)
			}
		})
	}
}
