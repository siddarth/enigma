package enigma

import (
	"math/rand"
	"strings"
	"time"
)

// The alphabet for this Enigma engine is (lowercase=only) a-z.
const alphabet = `abcdefghijklmnopqrstuvwxyz`

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Returns a re-ordered alphabet.
func randomAlphabet() string {
	src := strings.Split(alphabet, "")
	dest := make([]string, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}

	rand.Seed(rand.Int63())
	return strings.Join(dest, "")
}

// Returns 13 pairs of randomly-paired letters from the alphabet.
func randomAlphabetMapping() map[rune]rune {
	chars := randomAlphabet()
	set1, set2 := chars[0:13], chars[13:]
	mapping := make(map[rune]rune)

	for i, c := range set1 {
		v := rune(set2[i])
		mapping[c] = v
		mapping[v] = c
	}

	return mapping
}
