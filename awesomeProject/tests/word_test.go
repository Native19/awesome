package tests

import (
	"math/rand"
	"testing"
	"time"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error("IsPalindrome(\"detartrated\") = false")
	}
	if !IsPalindrome("kayak") {
		t.Error("IsPalindrome(\"kayak\") = false")
	}
}
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error("IsPalindrome(\"palindrome\") = true")
	}
}
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("ete") {
		t.Error("IsPalindrome(\"ete\") = false")
	}
}
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf("IsPalindrome(%q) = false", input)
	}
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"ete", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // He палиндром
		{"desserts", false},   // Полупалиндром
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// randomPalindrome возвращает палиндром, длина и содержимое
// которого задаются генератором псевдослучайных чисел rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // Случайная длина до 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x10)) // Случайная руна до '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Инициализация генератора псевдослучайных чисел,
	seed := time.Now().UTC().UnixNano()
	t.Logf("ГПСЧ инициализирован: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 100; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
