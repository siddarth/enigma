package enigma

import "testing"

func TestPlugboardTransform(t *testing.T) {
	type testCase struct {
		s map[rune]rune
		c rune
		r rune
	}

	testBank := []testCase{
		testCase{
			map[rune]rune{'a': 'b', 'c': 'd'}, 'a', 'b',
		},
		testCase{
			map[rune]rune{'a': 'b', 'c': 'd'}, 'e', 'e',
		},
		testCase{
			map[rune]rune{'a': 'b', 'c': 'd'}, 'c', 'd',
		},
	}

	for _, tc := range testBank {
		p := Plugboard{tc.s}
		r := p.Transform(tc.c)
		if r != tc.r {
			t.Errorf("invalid transform: expected %q, got %q", tc.r, r)
		}
	}
}
