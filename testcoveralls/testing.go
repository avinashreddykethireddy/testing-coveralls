package testcoveralls

func Add(x, y int) int {
	return x + y
}

// comment234
func CheckValues(x, y int) string {
	sum := x + y
	// helloo
	if sum > 10 {
		return "more"
	} else if sum == 0 {
		return "zero"
	} else {
		return "less"
	}
}
