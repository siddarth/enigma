package enigma

import (
	"log"
	"testing"
)

func TestEnigma(t *testing.T) {
	e := Enigma{}
	e.Init()

	plaintext := "helloworld"

	encrypted, err := e.Encode(plaintext)
	if err != nil {
		t.Fatalf("%+v", e)
	}
	log.Printf("Encrypted string: %q", encrypted)

	// Reset all the rotors to ignore ticks.
	for _, r := range e.Rotors {
		r.Reset()
	}

	decrypted, err := e.Encode(encrypted)
	if err != nil {
		t.Fatalf("%+v", e)
	}
	log.Printf("Decrypted string: %q", decrypted)

	if decrypted != plaintext {
		t.Errorf("incorrect decryption/encryption: %q", decrypted)
	}
}
