package salesbymatch

func sockMerchant(n int32, ar []int32) int32 {
	var result int32
	colorCounts := make(map[int32]int32)
	for _, c := range ar {
		colorCounts[c]++
	}
	for _, v := range colorCounts {
		result += v / 2
	}
	return result
}
