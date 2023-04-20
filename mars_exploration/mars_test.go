package mars

import "testing"

func TestMarsExploration(t *testing.T) {
	var tests = []struct {
		input string
		want  int32
	}{
		{"SOSSPSSQSSOR", 3},
		{"SOSSOT", 1},
		{"SOSSOSSOS", 0},
	}

	for _, tst := range tests {
		t.Run(tst.input, func(t *testing.T) {
			got := marsExploration(tst.input)
			if got != tst.want {
				t.Errorf("got %d, want %d", got, tst.want)
			}
		})
	}
}
