package catsandmouse

import (
	"fmt"
	"testing"
)

func TestCatsAndMouse(t *testing.T) {
	var tests = []struct {
		x    int32
		y    int32
		z    int32
		want string
	}{
		{1, 2, 3, "Cat B"},
		{1, 3, 2, "Mouse C"},
	}

	for _, tst := range tests {
		name := fmt.Sprintf("%d %d %d", tst.x, tst.y, tst.z)
		t.Run(name, func(t *testing.T) {
			got := catAndMouse(tst.x, tst.y, tst.z)
			if got != tst.want {
				t.Errorf("got %s, want %s", got, tst.want)
			}
		})
	}
}
