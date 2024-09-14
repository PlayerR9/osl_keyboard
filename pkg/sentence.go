package pkg

import (
	"bytes"
	"fmt"
	"strings"

	gcslc "github.com/PlayerR9/go-commons/slices"
)

// SentenceString returns the romanized sentence.
//
// Parameters:
//   - sentence: the sentence to romanize.
//
// Returns:
//   - string: the romanized sentence.
func SentenceString(sentence []*Syllable) string {
	if sentence == nil {
		return ""
	}

	var builder strings.Builder

	for _, syllable := range sentence {
		if syllable == nil {
			continue
		}

		builder.WriteString(syllable.String())
	}

	return builder.String()
}

// Sentence is a collection of syllables.
type Sentence struct {
	// syllables is the list of syllables in the sentence.
	syllables []*Syllable
}

// CheckValidity checks if the sentence is valid.
//
// Parameters:
//   - s: the sentence to check.
//
// Returns:
//   - error: nil if the sentence is valid, an error otherwise.
//
// Does nothing if the sentence is empty.
func CheckValidity(s []*Syllable) error {
	if len(s) == 0 {
		return nil
	}

	// 1. There cannot be more than three consonants in the same syllable.

	for i, syllable := range s {
		ok := syllable.check_consonant_count()
		if !ok {
			return fmt.Errorf("syllable at index %d has more than three consonants", i)
		}
	}

	// 2. There cannot be more than two consonants in a row.

	for i, syllable := range s {
		ok := syllable.check_consecutive_consonants()
		if !ok {
			return fmt.Errorf("syllable at index %d has more than two consonants in a row", i)
		}
	}

	// 3. There cannot be two tones in a row. Also, when a tone marker exists, there must exists either a
	// vowel or a consonant after it.

	tone := false
	var next_s *Syllable
	var err error

	for j, syllable := range s {
		if j+1 < len(s) {
			next_s = s[j+1]
		} else {
			next_s = nil
		}

		tone, err = syllable.check_tone(tone, next_s)
		if err != nil {
			return err
		}
	}

	return nil
}

// FinalTweaks applies changes to the sentence that are not part of the romanization process.
//
// Parameters:
//   - s: the sentence to fix.
//
// Does nothing if the sentence is empty.
func FinalTweaks(s []*Syllable) {
	if len(s) == 0 {
		return
	}

	// 1. If the character is the modifier 'a' and the next character is neither a vowel nor the last character,
	// then remove it.

	for _, syllable := range s {
		syllable.FinalTweaks()
	}

	// 2. If the current character is a tone marker, then change the following syllable to its lowercase variant.

	next_lowercase := false

	for _, syllable := range s {
		next_lowercase = syllable.fix_tone_variant(next_lowercase)
	}

	// If the current character is a modifier, then change its variant
	for _, syllable := range s {
		syllable.fix_relative_variants()
	}
}

// Tokenize returns the romanized sentence.
//
// Parameters:
//   - data: the data to tokenize.
//
// Returns:
//   - *Sentence: the romanized sentence.
//   - error: nil if the sentence is valid, an error otherwise.
func Tokenize(data []byte) ([]*Syllable, error) {
	var tokens []*Syllable

	syllable := NewSyllable()

	var prev *Character = nil

	for len(data) > 0 {
		// Remove leading and trailing spaces
		data = bytes.TrimSpace(data)

		// If the syllable is full, add it to the tokens
		if syllable.Size() == 3 {
			tokens = append(tokens, syllable)

			syllable = NewSyllable()
		}

		// Extract the first token
		solutions, err := ExtractFirstToken(data)
		if err != nil {
			return nil, err
		}

		// Debugging
		// Debugger.Println("N* of solutions:", len(solutions))
		//
		// for _, solution := range solutions {
		// 	Debugger.Println(solution.GetRomanization())
		// }
		// Debugger.Println()

		// Remove invalid tokens
		filterInvalidTokens := func(cd *CharacterDescription) bool {
			return cd.Type != CT_Coda ||
				modifierMatch(syllable, prev) ||
				helperMatch(syllable, prev)
		}

		solutions = gcslc.SliceFilter(solutions, filterInvalidTokens)

		// If no token was found, return an error
		if len(solutions) == 0 {
			// Debugger.Println("Syllables Done:")
			// Debugger.Println(tokens.String())
			//
			// Debugger.Println("Syllable in progress:")
			// Debugger.Print(syllable.RomanizedString())
			//
			// Debugger.Println()
			//
			// Debugger.Println("Remaining line:", line)

			return nil, fmt.Errorf("no token found in line: %s", data)
		}

		// matches
		if len(solutions) > 1 {
			panic(fmt.Sprintf("more than one match found for %s on characterSpace", data))
		}

		rom := solutions[0].Romanization
		data = data[len(rom):]

		c_type := solutions[0].Type

		switch c_type {
		case CT_Vowel:
			new_vowel, err := NewCharacter(solutions[0], 1)
			if err != nil {
				return nil, err
			}

			if prev == nil {
				new_vowel.SetVariant(0) // first vowel is uppercased
			}

			prev = new_vowel

			syllable.Append(new_vowel)
		case CT_Consonant:
			new_consonant, err := NewCharacter(solutions[0], 1)
			if err != nil {
				return nil, err
			}

			if prev == nil {
				new_consonant.SetVariant(0) // first consonant is uppercased
			}

			prev = new_consonant

			if syllable.Size() != 0 {
				tokens = append(tokens, syllable)

				syllable = NewSyllable()
			}

			syllable.Append(new_consonant)
		case CT_Coda:
			new_modifier, err := NewCharacter(solutions[0], 0)
			if err != nil {
				return nil, err
			}

			prev = new_modifier

			syllable.Append(new_modifier)
		case CT_Extra:
			prev = nil

			if syllable.Size() != 0 {
				tokens = append(tokens, syllable)

				syllable = NewSyllable()
			}

			c, err := NewCharacter(solutions[0], 0)
			if err != nil {
				return nil, err
			}

			syllable.Append(c)

			tokens = append(tokens, syllable)

			syllable = NewSyllable()
		case CT_Helper:
			prev = nil

			if syllable.Size() != 0 {
				tokens = append(tokens, syllable)

				syllable = NewSyllable()
			}

			if solutions[0].Romanization != "|" {
				c, err := NewCharacter(solutions[0], 0)
				if err != nil {
					return nil, err
				}

				syllable.Append(c)

				tokens = append(tokens, syllable)

				syllable = NewSyllable()
			}
		default:
			panic(fmt.Sprintf("unknown character type: %d", c_type))
		}
	}

	// If there is a syllable in progress, add it to the tokens
	if syllable.Size() != 0 {
		tokens = append(tokens, syllable)
	}

	return tokens, nil
}
