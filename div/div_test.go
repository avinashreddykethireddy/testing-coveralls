package div

import "testing"

func TestDiv(t *testing.T) {
	sum := Div(2, 2)
	expected := 0

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestCheckDiv(t *testing.T) {
	cases := []struct {
		x        int
		y        int
		expected string
	}{
		{10, 6, "more"},
		{20, 10, "zero"},
	}
	for _, c := range cases {
		t.Run("testing", func(t *testing.T) {
			got := CheckDiv(c.x, c.y)
			if got != c.expected {
				t.Fatalf("Wanted %v , but got %v", c.expected, got)
			}
		})
	}
}
