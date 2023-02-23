package subber

func Sub(x, y int) int {
	return x - y
}

// comment234
func CheckSub(x, y int) string {
	sum := x - y
	if sum > 0 {
		return "more"
	} else if sum == 0 {
		return "zero"
	} else {
		return "less"
	}
}
