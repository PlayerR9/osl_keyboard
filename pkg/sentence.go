package pkg

import (
	"fmt"
	"strings"

	dba "github.com/PlayerR9/go-debug/assert"
	prx "github.com/PlayerR9/osl_keyboard/pkg/parsing"
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
	tokens, err := prx.FullLex(data)
	if err != nil {
		return nil, err
	}

	var syllables []*Syllable

	var builder SyllableBuilder

	var prev *Character = nil

	for i := 0; i < len(tokens)-1; i++ {
		// 1. If the syllable is "full", then add it to the list of syllables.
		if builder.Size() == 3 {
			syllables = append(syllables, builder.Build())
			builder.Reset()
		}

		// TODO: Handle the newline character

		// 2. Ensure that is not an invalid combination.
		sol, ok := DescriptionFromRomanization(tokens[i].Data)
		dba.AssertOk(ok, "DescriptionFromRomanization(%q)", tokens[i].Data)

		if sol.Type == CT_Coda && !builder.modifierMatch(prev) && !builder.helperMatch(prev) {
			return nil, fmt.Errorf("unmatched coda")
		}

		c_type := sol.Type

		switch c_type {
		case CT_Vowel:
			new_vowel, err := NewCharacter(sol, 1)
			if err != nil {
				return nil, err
			}

			if prev == nil {
				new_vowel.SetVariant(0) // first vowel is uppercased
			}

			prev = new_vowel

			builder.Append(new_vowel)
		case CT_Consonant:
			new_consonant, err := NewCharacter(sol, 1)
			if err != nil {
				return nil, err
			}

			if prev == nil {
				new_consonant.SetVariant(0) // first consonant is uppercased
			}

			prev = new_consonant

			if builder.Size() != 0 {
				syllables = append(syllables, builder.Build())
				builder.Reset()
			}

			builder.Append(new_consonant)
		case CT_Coda:
			new_modifier, err := NewCharacter(sol, 0)
			if err != nil {
				return nil, err
			}

			prev = new_modifier

			builder.Append(new_modifier)
		case CT_Extra:
			prev = nil

			if builder.Size() != 0 {
				syllables = append(syllables, builder.Build())
				builder.Reset()
			}

			c, err := NewCharacter(sol, 0)
			if err != nil {
				return nil, err
			}

			builder.Append(c)

			syllables = append(syllables, builder.Build())
			builder.Reset()
		case CT_Helper:
			prev = nil

			if builder.Size() != 0 {
				syllables = append(syllables, builder.Build())
				builder.Reset()
			}

			if sol.Romanization != "|" {
				c, err := NewCharacter(sol, 0)
				if err != nil {
					return nil, err
				}

				builder.Append(c)

				syllables = append(syllables, builder.Build())
				builder.Reset()
			}
		default:
			return nil, fmt.Errorf("unknown character type: %d", c_type)
		}
	}

	// If there is a syllable in progress, add it to the tokens
	if builder.Size() != 0 {
		syllables = append(syllables, builder.Build())
	}

	return syllables, nil
}
