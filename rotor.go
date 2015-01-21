package enigma

import (
	"fmt"
	"strings"
)

// Represents a simplified version of the Enigma rotor, as
// described here:
// http://en.wikipedia.org/wiki/Enigma_rotor_details
type Rotor struct {
	// This is the ordering of characters for this specific Rotor.
	Chars string
	// This keeps track of the initial state of Chars: calls to
	// Tick() mutates Chars, so this is for convenience to reset
	// the state of the rotor.
	initState string
}

// Initializes a random ordering for the characters on the rotor.
func (r *Rotor) Init() {
	chars := randomAlphabet()
	r.initState = chars
	r.Chars = chars
}

// Resets Chars to initState (as set up in Init()): used as a
// convenience method for testing.
func (r *Rotor) Reset() {
	r.Chars = r.initState
}

// Simulates the flow of current through this rotor after the
// reflector. See the image on this page for more information:
// http://enigma.louisedade.co.uk/howitworks.html
func (r *Rotor) Reflect(c rune) (rune, error) {
	return r.process(c, r.Chars, alphabet)
}

// Simulates the flow of current through this rotor before it
// reaches the reflector.
func (r *Rotor) Transform(c rune) (rune, error) {
	return r.process(c, alphabet, r.Chars)
}

// Helper method used to implement Transform and Reflect.
func (r *Rotor) process(c rune, lookup string, substitution string) (rune, error) {
	pos := strings.IndexRune(lookup, c)
	if pos == -1 {
		return 0, fmt.Errorf("character outside of alphabet: %q", c)
	}

	return rune(substitution[pos]), nil
}

// This represents the actual movement of the rotor: an important
// property of Enigma is that multiple occurrences of the same
// letter in a string results in different corresponding letters
// in the encoded string. This is accomplished by one of the rotors
// moving for each keypress.
func (r *Rotor) Tick() {
	r.Chars = r.Chars[1:] + string(r.Chars[0])
}
