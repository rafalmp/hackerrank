package countingvalleys

func countingValleys(steps int32, path string) int32 {
	var level, previous, result int32
	for _, c := range path {
		if c == 'D' {
			level--
			if level < 0 && previous == 0 {
				result++
			}
		} else {
			level++
		}
		previous = level
	}
	return result
}
