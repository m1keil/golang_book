// Modify randomPalindrome to exercise IsPalindrome’s handling of punctuation and spaces.
package ex4

import (
	"math/rand"
	"testing"
	"time"
	"unicode"

	word "gopl.io/ch11/word2"
)

// build pool of letters, punctuation and space runes
var runePool = getPool()

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := runePool[rng.Intn(len(runePool))]
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func getPool() (output []rune) {
	for _, r := range unicode.Space.R16 {
		for c := r.Lo; c <= r.Hi; c += r.Stride {
			output = append(output, rune(c))
		}
	}

	for _, r := range unicode.Punct.R16 {
		for c := r.Lo; c <= r.Hi; c += r.Stride {
			output = append(output, rune(c))
		}
	}

	for _, r := range unicode.Letter.R16[:20] {
		for c := r.Lo; c <= r.Hi; c += r.Stride {
			output = append(output, rune(c))
		}
	}

	return output
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
