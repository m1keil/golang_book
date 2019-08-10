// TestRandomPalindromes only tests palindromes. Write a randomized test that generates and verifies non-palindromes.
package ex3

import (
	"gopl.io/ch11/word2"
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomNonPalindrome(rng *rand.Rand) string {
	used := make(map[rune]bool)
	n := rng.Intn(25) + 2 // random length between 2 to 24 (avoid 1 char strings)
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		var r rune
		for {
			r = rune(rng.Intn(0x01000)) // random rune up to '\u0999
			// only allow letters and each letter allowed only once
			if unicode.IsLetter(r) && !used[unicode.ToLower(r)] { break }
		}
		used[r] = true
		runes[i] = r
	}
	return string(runes)
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}