package hurdlerace

import (
	"fmt"
	"testing"
)

func TestHurdleRace(t *testing.T) {
	var tests = []struct {
		k      int32
		height []int32
		want   int32
	}{
		{1, []int32{1, 2, 3, 3, 2}, 2},
		{4, []int32{1, 6, 3, 5, 2}, 2},
		{7, []int32{2, 5, 4, 5, 2}, 0},
	}

	for i, tst := range tests {
		name := fmt.Sprintf("Case %d", i+1)
		t.Run(name, func(t *testing.T) {
			got := hurdleRace(tst.k, tst.height)
			if got != tst.want {
				t.Errorf("got %d, want %d", got, tst.want)
			}
		})
	}
}
