package electronicsshop

import "testing"

func TestCountingValleys(t *testing.T) {
	var tests = []struct {
		name      string
		keyboards []int32
		drives    []int32
		budget    int32
		want      int32
	}{
		{"Sample input 1", []int32{3, 1}, []int32{5, 2, 8}, 10, 9},
		{"Sample input 2", []int32{4}, []int32{5}, 5, -1},
		{"Sample input 3", []int32{4, 5, 6}, []int32{5, 6, 7}, 12, 12},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := getMoneySpent(tst.keyboards, tst.drives, tst.budget)
			if got != tst.want {
				t.Errorf("got %d, want %d", got, tst.want)
			}
		})
	}
}
