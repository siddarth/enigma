package enigma

// Represents a simplified version of the Enigma plugboard, as
// described here:
// http://en.wikipedia.org/wiki/Enigma_machine#Plugboard
type Plugboard struct {
	// The "steckering" is just a mapping of characters from
	// one to another.
	Steckering map[rune]rune
}

// Initialize a random a fully Steckered set of character
// mappings.
func (p *Plugboard) Init() {
	p.Steckering = randomAlphabetMapping()
}

// Return the mapping for a given character. The Plugboard
// is quite literally just a dictionary, with the difference
// that if a given character is not steckered, the same
// character is returned back to the engine.
func (p *Plugboard) Transform(c rune) rune {
	if sub, ok := p.Steckering[c]; !ok {
		return c
	} else {
		return sub
	}
}
