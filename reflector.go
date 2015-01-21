package enigma

import "fmt"

// Represents a reflector as described here:
// http://en.wikipedia.org/wiki/Enigma_machine#Reflector
// An important property of the reflector that is maintained
// in this simulator is that no character maps to itself.
type Reflector struct {
	Mapping map[rune]rune
}

// Initializes the Reflector with some random mapping.
func (r *Reflector) Init() {
	r.Mapping = randomAlphabetMapping()
}

// Reflect the given character based on the mapping.
// Note that while Reflectors and Plugboards are similar,
// an important difference is that a plugboard returns
// the input character when transforming an unsteckered
// character. The Reflector, on the other hand, *must* keep
// track of all 13 pairs of characters.
func (r *Reflector) Transform(c rune) (rune, error) {
	if sub, ok := r.Mapping[c]; !ok {
		return 0, fmt.Errorf("unknown character: %q", c)
	} else {
		return sub, nil
	}
}
