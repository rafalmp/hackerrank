package mars

func marsExploration(s string) int32 {
	var msg = []rune{'S', 'O', 'S'}
	var result int32 = 0
	for i, c := range s {
		if c != msg[i%3] {
			result++
		}
	}
	return result
}
