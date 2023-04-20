package caesar

import (
	"testing"
)

func TestCaesar(t *testing.T) {
	var tests = []struct {
		input  string
		offset int32
		want   string
	}{
		{"There's-a-starman-waiting-in-the-sky", 3, "Wkhuh'v-d-vwdupdq-zdlwlqj-lq-wkh-vnb"},
		{"middle-Outz", 2, "okffng-Qwvb"},
	}

	for _, tst := range tests {
		t.Run(tst.input, func(t *testing.T) {
			got := caesarCipher(tst.input, tst.offset)
			if got != tst.want {
				t.Errorf("got %s, want %s", got, tst.want)
			}
		})
	}
}
