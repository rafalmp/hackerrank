package hrinastring

func hackerrankInString(s string) string {
	var needle = "hackerrank"
	var from = 0
nextchar:
	for _, nc := range needle {
		for _, sc := range s[from:] {
			from++
			if sc == nc {
				continue nextchar
			}
		}
		return "NO"
	}
	return "YES"
}
