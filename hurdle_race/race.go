package hurdlerace

func hurdleRace(k int32, height []int32) (result int32) {
	for _, h := range height {
		if d := h - k; d > result {
			result = d
		}
	}
	return
}
