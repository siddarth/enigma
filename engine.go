// Package enigma is an Enigma engine simulator.
// For further reading:
//
// http://en.wikipedia.org/wiki/Enigma_machine
//
// http://enigma.louisedade.co.uk/howitworks.html
package enigma

const numRotors = 3

// The core Enigma simulation engine.
type Enigma struct {
	// Represents a plugboard as described here:
	// http://en.wikipedia.org/wiki/Enigma_machine#Plugboard
	*Plugboard

	// Represents a rotor as described here:
	// http://en.wikipedia.org/wiki/Enigma_rotor_details
	Rotors []*Rotor

	// Represents a reflector as described here:
	// http://en.wikipedia.org/wiki/Enigma_machine#Reflector
	Reflector *Reflector
}

// Initializes the Enigma engine. This involves setting up each
// of its components (Plugboard, Rotors, and Reflectors), and
// initializing them with random alphabets and ordering.
func (e *Enigma) Init() {
	pb := &Plugboard{}
	pb.Init()
	e.Plugboard = pb

	rf := &Reflector{}
	rf.Init()
	e.Reflector = rf

	rotors := make([]*Rotor, numRotors)
	for i := 0; i < 3; i++ {
		r := &Rotor{}
		r.Init()

		rotors[i] = r
	}
	e.Rotors = rotors
}

// This is the main API method to encode a given string using
// the Enigma simulator.
// http://en.wikipedia.org/wiki/Enigma_machine
// http://enigma.louisedade.co.uk/howitworks.html
func (e *Enigma) Encode(str string) (string, error) {
	encoded := ""

	for n, c := range str {
		x := c
		var err error

		// Each keypress first passes through the plugboard,
		// resulting in possible steckering.
		x = e.Plugboard.Transform(x)

		// The steckered characters are repeatedly passed through
		// each of the rotors.
		if x, err = e.rotorTransform(x); err != nil {
			return "", err
		}

		// The reflector provides yet another transformation at
		// the end of the rotor-passage.
		if x, err = e.Reflector.Transform(x); err != nil {
			return "", err
		}

		// Now, current passes back through the rotors in the
		// opposite direction.
		if x, err = e.rotorReflect(x); err != nil {
			return "", err
		}

		// Another transformation through the plugboard occurs.
		x = e.Plugboard.Transform(x)

		// One of the rotors in the machine rotates once.
		e.Rotors[(n/26)%numRotors].Tick()

		encoded += string(x)
	}

	return encoded, nil
}

// Helper method to implement the outgoing (to the reflector)
// transformations through the rotors.
func (e *Enigma) rotorTransform(c rune) (rune, error) {
	x := c
	var err error

	for i := 0; i < len(e.Rotors); i++ {
		r := e.Rotors[i]
		x, err = r.Transform(x)

		if err != nil {
			return 0, err
		}
	}

	return x, nil
}

// Helper method to implement the transformations through the
// rotors after being reflected.
func (e *Enigma) rotorReflect(c rune) (rune, error) {
	x := c
	var err error

	for i := len(e.Rotors) - 1; i >= 0; i-- {
		r := e.Rotors[i]
		x, err = r.Reflect(x)

		if err != nil {
			return 0, err
		}
	}

	return x, nil
}
