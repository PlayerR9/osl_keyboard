package pkg

import (
	"bytes"

	gcch "github.com/PlayerR9/go-commons/runes"
)

// character_table is a table of all character descriptions.
var character_table []*CharacterDescription

func init() {
	character_table = make([]*CharacterDescription, 0, 28)

	// Vowels
	character_table = append(character_table, new_character_description("u", CT_Vowel, CQ_Narrow, '\uE001', '\uE01B'))
	character_table = append(character_table, new_character_description("ö", CT_Vowel, CQ_Narrow, '\uE004', '\uE01E'))
	character_table = append(character_table, new_character_description("y", CT_Vowel, CQ_Narrow, '\uE006', '\uE020'))
	character_table = append(character_table, new_character_description("e", CT_Vowel, CQ_Wide, '\uE00C', '\uE026'))
	character_table = append(character_table, new_character_description("i", CT_Vowel, CQ_Narrow, '\uE00F', '\uE029'))
	character_table = append(character_table, new_character_description("a", CT_Vowel, CQ_Narrow, '\uE012', '\uE02C'))
	character_table = append(character_table, new_character_description("o", CT_Vowel, CQ_Wide, '\uE015', '\uE02F'))
	character_table = append(character_table, new_character_description("ë", CT_Vowel, CQ_Narrow, '\uE017', '\uE031'))
	character_table = append(character_table, new_character_description("ä", CT_Vowel, CQ_Narrow, '\uE009', '\uE023', '\uE036', '\uE03B'))

	// Consonants
	character_table = append(character_table, new_character_description("r", CT_Consonant, CQ_Wide, '\uE000', '\uE01A'))
	character_table = append(character_table, new_character_description("s", CT_Consonant, CQ_Wide, '\uE002', '\uE01C'))
	character_table = append(character_table, new_character_description("v", CT_Consonant, CQ_Wide, '\uE003', '\uE01D'))
	character_table = append(character_table, new_character_description("m", CT_Consonant, CQ_Narrow, '\uE005', '\uE01F'))
	character_table = append(character_table, new_character_description("k", CT_Consonant, CQ_Narrow, '\uE008', '\uE022'))
	character_table = append(character_table, new_character_description("b", CT_Consonant, CQ_Narrow, '\uE00B', '\uE025'))
	character_table = append(character_table, new_character_description("n", CT_Consonant, CQ_Narrow, '\uE014', '\uE02E'))
	character_table = append(character_table, new_character_description("t", CT_Consonant, CQ_Narrow, '\uE016', '\uE030'))
	character_table = append(character_table, new_character_description("d", CT_Consonant, CQ_Narrow, '\uE018', '\uE032'))

	// Modifiers
	character_table = append(character_table, new_character_description("ṇ", CT_Coda, CQ_Narrow, '\uE037', '\uE03C'))
	character_table = append(character_table, new_character_description("ṃ", CT_Coda, CQ_Narrow, '\uE038', '\uE03D'))
	character_table = append(character_table, new_character_description("l", CT_Coda, CQ_Narrow, '\uE039', '\uE03E'))

	// Extras
	character_table = append(character_table, new_character_description(".", CT_Extra, CQ_Narrow, '\uE034'))

	// Tone markers
	character_table = append(character_table, new_character_description("\\m", CT_Extra, CQ_Narrow, '\uE04C'))
	character_table = append(character_table, new_character_description("\\h", CT_Extra, CQ_Narrow, '\uE04D'))
	character_table = append(character_table, new_character_description("\\r", CT_Extra, CQ_Narrow, '\uE04E'))
	character_table = append(character_table, new_character_description("\\l", CT_Extra, CQ_Narrow, '\uE04F'))

	// Helpers
	character_table = append(character_table, new_character_description("|", CT_Helper, CQ_Narrow))
	character_table = append(character_table, new_character_description(".", CT_Helper, CQ_Narrow, '\uE034'))
}

// Alphabet returns a table of runes representing the entire alphabet.
//
// Returns:
//   - [][]rune: the rune table.
func Alphabet() [][]rune {
	var alphabet [][]rune

	for _, char := range character_table {
		alphabet = append(alphabet, char.variants)
	}

	return alphabet
}

// ExtractFirstToken extracts the first token from the input bytes.
//
// Parameters:
//   - input: the input bytes.
//
// Returns:
//   - []*CharacterDescription: the first token extracted from the input bytes.
//   - error: an error if the input is invalid.
func ExtractFirstToken(input []byte) ([]*CharacterDescription, error) {
	var solutions []*CharacterDescription

	// check if the input has a character at the beginning of the string by adding to the solutions the index of the longest match
	// if there are multiple matches of the same length, add all of them
	var y []rune

	maxIndex := -1

	for i, char := range character_table {
		// convert the character to utf8
		x, err := gcch.StringToUtf8(char.Romanization)
		if err != nil {
			return nil, err
		}

		// check if the input starts with the character
		if !bytes.HasPrefix(input, []byte(char.Romanization)) {
			continue
		}

		// check if the character is longer than the previous one or if it is the first one
		// if it is, change the max_index
		if maxIndex == -1 || len(x) > len(y) {
			maxIndex = i
			y = x

			// reset the solutions
			solutions = []*CharacterDescription{char}
		} else if len(x) == len(y) { // if the character is the same length as the previous one, add it to the solutions
			solutions = append(solutions, char)
		}
	}

	return solutions, nil
}
