package pickingnumbers

import (
	"fmt"
	"testing"
)

func TestPickingNumbers(t *testing.T) {
	var tests = []struct {
		a    []int32
		want int32
	}{
		{[]int32{1, 1, 2, 2, 4, 4, 5, 5, 5}, 5},
		{[]int32{4, 6, 5, 3, 3, 1}, 3},
		{[]int32{1, 2, 2, 3, 1, 2}, 5},
		{[]int32{66, 66, 66, 66, 66}, 5},
		{[]int32{4, 97, 5, 97, 99, 97, 4, 97, 5, 97}, 5},
	}

	for i, tst := range tests {
		name := fmt.Sprintf("Case %d", i+1)
		t.Run(name, func(t *testing.T) {
			got := pickingNumbers(tst.a)
			if got != tst.want {
				t.Errorf("got %d, want %d", got, tst.want)
			}
		})
	}
}
