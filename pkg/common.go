package pkg

import (
	"fmt"
	"strings"

	sext "github.com/PlayerR9/MyGoLib/Utility/StringExt"
)

// TO DO: add space as a meaningful character

var characterSpace []*CharacterDescription

func init() {
	characterSpace = []*CharacterDescription{
		// Vowels
		NewCharacterDescription("u", CT_Vowel, CQ_Narrow, '\uE001', '\uE01B'),
		NewCharacterDescription("ö", CT_Vowel, CQ_Narrow, '\uE004', '\uE01E'),
		NewCharacterDescription("y", CT_Vowel, CQ_Narrow, '\uE006', '\uE020'),
		NewCharacterDescription("e", CT_Vowel, CQ_Wide, '\uE00C', '\uE026'),
		NewCharacterDescription("i", CT_Vowel, CQ_Narrow, '\uE00F', '\uE029'),
		NewCharacterDescription("a", CT_Vowel, CQ_Narrow, '\uE012', '\uE02C'),
		NewCharacterDescription("o", CT_Vowel, CQ_Wide, '\uE015', '\uE02F'),
		NewCharacterDescription("ë", CT_Vowel, CQ_Narrow, '\uE017', '\uE031'),
		NewCharacterDescription("ä", CT_Vowel, CQ_Narrow, '\uE009', '\uE023', '\uE036', '\uE03B'),

		// Consonants
		NewCharacterDescription("r", CT_Consonant, CQ_Wide, '\uE000', '\uE01A'),
		NewCharacterDescription("s", CT_Consonant, CQ_Wide, '\uE002', '\uE01C'),
		NewCharacterDescription("v", CT_Consonant, CQ_Wide, '\uE003', '\uE01D'),
		NewCharacterDescription("m", CT_Consonant, CQ_Narrow, '\uE005', '\uE01F'),
		NewCharacterDescription("k", CT_Consonant, CQ_Narrow, '\uE008', '\uE022'),
		NewCharacterDescription("b", CT_Consonant, CQ_Narrow, '\uE00B', '\uE025'),
		NewCharacterDescription("n", CT_Consonant, CQ_Narrow, '\uE014', '\uE02E'),
		NewCharacterDescription("t", CT_Consonant, CQ_Narrow, '\uE016', '\uE030'),
		NewCharacterDescription("d", CT_Consonant, CQ_Narrow, '\uE018', '\uE032'),

		// Modifiers
		NewCharacterDescription("ṇ", CT_Coda, CQ_Narrow, '\uE037', '\uE03C'),
		NewCharacterDescription("ṃ", CT_Coda, CQ_Narrow, '\uE038', '\uE03D'),
		NewCharacterDescription("l", CT_Coda, CQ_Narrow, '\uE039', '\uE03E'),

		// Extras
		NewCharacterDescription(".", CT_Extra, CQ_Narrow, '\uE034'),

		// Tone markers
		NewCharacterDescription("\\m", CT_Extra, CQ_Narrow, '\uE04C'),
		NewCharacterDescription("\\h", CT_Extra, CQ_Narrow, '\uE04D'),
		NewCharacterDescription("\\r", CT_Extra, CQ_Narrow, '\uE04E'),
		NewCharacterDescription("\\l", CT_Extra, CQ_Narrow, '\uE04F'),

		// Helpers
		NewCharacterDescription("|", CT_Helper, CQ_Narrow),
		NewCharacterDescription(".", CT_Helper, CQ_Narrow, '\uE034'),
	}
}

func GetAlphabet() [][]rune {
	// initialize the return value
	alphabet := make([][]rune, 0)

	for _, char := range characterSpace {
		alphabet = append(alphabet, char.variants)
	}

	return alphabet
}

/*
	ExtractFirstToken extracts the first token from the input string

Parameters:

	input: the input string

Returns:

	[]int: the first token extracted from the input string
*/
func ExtractFirstToken(input string) []*CharacterDescription {
	// initialize the return value
	solutions := make([]*CharacterDescription, 0)

	// check if the input has a character at the beginning of the string by adding to the solutions the index of the longest match
	// if there are multiple matches of the same length, add all of them
	var y []rune

	maxIndex := -1

	for i, char := range characterSpace {
		// convert the character to utf8
		x, err := sext.ToUTF8Runes(char.romanization)
		if err != nil {
			panic(fmt.Errorf("error converting %s to utf8: %s", char.romanization, err.Error()))
		}

		// check if the input starts with the character
		if !strings.HasPrefix(input, char.romanization) {
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

	return solutions
}
