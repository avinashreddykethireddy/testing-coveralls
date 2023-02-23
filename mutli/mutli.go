package mutli

func Mutli(x, y int) int {
	return x * y
}

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
