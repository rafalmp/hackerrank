package electronicsshop

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var result int32 = -1
	for _, k := range keyboards {
		for _, d := range drives {
			if sum := k + d; sum <= b && sum > result {
				result = sum
			}
		}
	}
	return result
}
