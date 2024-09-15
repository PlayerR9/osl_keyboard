package pkg

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
	character_table = append(character_table, new_character_description(".", CT_Extra, CQ_Narrow, '\uE034')) // duplicate: fix this

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

// DescriptionFromRomanization returns the description of a character from its romanization.
//
// Parameters:
//   - romanization: the romanization of the character.
//
// Returns:
//   - *CharacterDescription: the description of the character.
//   - bool: true if the character was found, false otherwise.
func DescriptionFromRomanization(romanization string) (*CharacterDescription, bool) {
	for _, char := range character_table {
		if char.Romanization == romanization {
			return char, true
		}
	}

	return nil, false
}
