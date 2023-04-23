package salesbymatch

import (
	"fmt"
	"testing"
)

func TestSockMerchant(t *testing.T) {
	var tests = []struct {
		n    int32
		ar   []int32
		want int32
	}{
		{7, []int32{1, 2, 1, 2, 1, 3, 2}, 2},
		{9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
	}

	for _, tst := range tests {
		name := fmt.Sprintf("%d socks", tst.n)
		t.Run(name, func(t *testing.T) {
			got := sockMerchant(tst.n, tst.ar)
			if got != tst.want {
				t.Errorf("got %d, want %d", got, tst.want)
			}
		})
	}
}
