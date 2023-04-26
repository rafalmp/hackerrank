package catsandmouse

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func catAndMouse(x int32, y int32, z int32) string {
	ac := abs(x - z)
	bc := abs(y - z)

	if ac < bc {
		return "Cat A"
	} else if ac > bc {
		return "Cat B"
	}
	return "Mouse C"
}
