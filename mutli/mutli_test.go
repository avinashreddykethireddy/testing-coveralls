package mutli

import "testing"

func TestMutli(t *testing.T) {
	sum := Mutli(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestCheckmulti(t *testing.T) {
	cases := []struct {
		x        int
		y        int
		expected string
	}{
		{2, -5, "less"},
		{20, 10, "more"},
		{20, 0, "zero"},
	}
	for _, c := range cases {
		t.Run("testing", func(t *testing.T) {
			got := Checkmulti(c.x, c.y)
			if got != c.expected {
				t.Fatalf("Wanted %v , but got %v", c.expected, got)
			}
		})
	}
}
