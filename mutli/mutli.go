package mutli

func Mutli(x, y int) int {
	return x * y
}

// comment234
func Checkmulti(x, y int) string {
	sum := x * y
	if sum > 0 {
		return "more"
	} else if sum == 0 {
		return "zero"
	} else {
		return "less"
	}
}
