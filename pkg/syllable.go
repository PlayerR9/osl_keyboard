package pkg

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// Syllable is a collection of characters.
type Syllable struct {
	// chars is the list of characters in the syllable.
	chars []*Character
}

// String returns the romanized syllable.
//
// Returns:
//   - string: the romanized syllable.
func (s Syllable) String() string {
	var builder strings.Builder

	for _, char := range s.chars {
		rom := char.GetRomanization()

		builder.WriteString(rom)
	}

	return builder.String()
}

// NewSyllable creates a new syllable.
//
// Returns:
//   - *Syllable: the new syllable. Never returns nil.
func NewSyllable() *Syllable {
	return &Syllable{
		chars: make([]*Character, 0),
	}
}

// Size returns the size of the syllable.
//
// Returns:
//   - int: the size of the syllable.
func (s Syllable) Size() int {
	return len(s.chars)
}

// Append adds a character to the syllable.
//
// Parameters:
//   - char: the character to add to the syllable.
//
// Does nothing if the receiver or the character are nil.
func (s *Syllable) Append(char *Character) {
	if s == nil || char == nil {
		return
	}

	s.chars = append(s.chars, char)
}

func (s *Syllable) check_consonant_count() bool {
	consonants := 0

	for _, char := range s.chars {
		charType := char.GetType()

		if charType == CT_Consonant {
			if consonants == 3 {
				return false
			}

			consonants++
		}
	}

	return true
}

func (s *Syllable) check_consecutive_consonants() bool {
	consonants := 0

	for _, char := range s.chars {
		charType := char.GetType()

		if charType == CT_Consonant {
			if consonants == 2 {
				return false
			}

			consonants++
		} else {
			consonants = 0
		}
	}

	return true
}

func (s *Syllable) check_tone(tone bool, nextS *Syllable) (bool, error) {
	for _, char := range s.chars {
		if !strings.HasPrefix(char.GetRomanization(), "\\") {
			tone = false
			continue
		}

		if nextS == nil {
			return tone, fmt.Errorf("a tone marker can't be the last character of a syllable")
		}

		next := nextS.chars[0]

		if next.GetType() != CT_Vowel && next.GetType() != CT_Consonant {
			return tone, errors.New("missing consonant or vowel after the tone marker")
		}

		if tone {
			return tone, errors.New("syllable has two tones in a row")
		}

		tone = true
	}

	return tone, nil
}

func (s *Syllable) FinalTweaks() {
	// if the character is a modifier 'a' and the next character is not a vowel or it
	// is the last character, remove it
	var prevC *Character = nil

	for i := 0; i < len(s.chars); i++ {
		currC := s.chars[i]

		rom := currC.GetRomanization()

		if rom != "ä" {
			if i == 0 {
				prevC = nil
			} else {
				prevC = currC
			}

			continue
		}

		found := false

		if i != 0 {
			prevCT := prevC.GetType()

			switch prevCT {
			case CT_Consonant:
				charType := s.chars[i+1].GetType()

				if i+1 < len(s.chars) && charType != CT_Coda {
					// if the character follows an onset consonant and the next character
					// is not a modifier, remove it
					s.chars = slices.Delete(s.chars, i, i+1)

					found = true
				}
			case CT_Vowel:
				// print the modifier version of it
				if prevC.IsNarrow() {
					currC.SetVariant(2)
				} else {
					currC.SetVariant(3)
				}

				found = true
			}
		}

		if !found {
			// if previous case didn't happen, print the non-modifier version of it
			currC.SetVariant(1)

			if i == 0 {
				currC.SetVariant(0) // first vowel is uppercased
			}
		}

		prevC = currC
	}
}

func (s *Syllable) fix_tone_variant(nextLowercase bool) bool {
	rom := s.chars[0].GetRomanization()

	if strings.HasPrefix(rom, "\\") {
		nextLowercase = true
	} else {
		if nextLowercase {
			s.chars[0].SetVariant(1)
		}

		nextLowercase = false
	}

	return nextLowercase
}

func (s *Syllable) fix_relative_variants() {
	for i := 0; i < len(s.chars); i++ {
		charType := s.chars[i].GetType()
		if charType != CT_Coda || i == 0 {
			continue
		}

		// Debugger.Println("Modifier found:", s.chars[i].String())
		// Debugger.Println()

		prev := s.chars[i-1]

		if prev.IsNarrow() {
			s.chars[i].SetVariant(0)
		} else {
			s.chars[i].SetVariant(1)
		}
	}
}
