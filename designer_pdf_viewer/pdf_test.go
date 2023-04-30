package designerpdfviewer

import "testing"

func TestPickingNumbers(t *testing.T) {
	var tests = []struct {
		h    []int32
		word string
		want int32
	}{
		{[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, "abc", 9},
		{[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}, "zaba", 28},
	}

	for _, tst := range tests {
		t.Run(tst.word, func(t *testing.T) {
			got := designerPdfViewer(tst.h, tst.word)
			if got != tst.want {
				t.Errorf("got %d, want %d", got, tst.want)
			}
		})
	}
}
