package div

func Div(x, y int) int {
	return x % y
}

// comment2
func CheckDiv(x, y int) string {
	sum := x % y
	if sum > 0 {
		return "more"
	} else {
		return "zero"
	}
}
