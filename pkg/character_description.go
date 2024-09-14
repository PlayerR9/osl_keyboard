package pkg

import "unicode/utf8"

// CharacterType represents the type of the character.
type CharacterType int

const (
	CT_Vowel CharacterType = iota
	CT_Consonant
	CT_Coda
	CT_Extra
	CT_Helper
)

// CharacterNarrowness represents the narrowness of the character.
type CharacterNarrowness int

const (
	CQ_Narrow CharacterNarrowness = iota
	CQ_Wide
)

// CharacterDescription helds the description of a character.
type CharacterDescription struct {
	// Romanization is the Romanization of the character.
	Romanization string

	// len is the length of the romanization. This is used for optimization purposes.
	len int

	// Type is the type of the character.
	Type CharacterType

	// narrowness is the narrowness of the character.
	narrowness CharacterNarrowness

	// variants is the list of variants of the character.
	variants []rune
}

// new_character_description creates a new CharacterDescription.
//
// Parameters:
//   - romanization: the romanization of the character.
//   - c_type: the type of the character.
//   - narrowness: the narrowness of the character.
//   - variants: the list of variants of the character.
//
// Returns:
//   - *CharacterDescription: the new CharacterDescription. Never returns nil.
func new_character_description(romanization string, c_type CharacterType, narrowness CharacterNarrowness, variants ...rune) *CharacterDescription {
	return &CharacterDescription{
		Romanization: romanization,
		len:          utf8.RuneCountInString(romanization),
		Type:         c_type,
		narrowness:   narrowness,
		variants:     variants,
	}
}

// VariantSize returns the number of variants of the character.
//
// Returns:
//   - int: the number of variants of the character.
func (cd CharacterDescription) VariantSize() int {
	return len(cd.variants)
}

// IsNarrow checks if the character is narrow.
//
// Returns:
//   - bool: true if the character is narrow, false otherwise.
func (cd CharacterDescription) IsNarrow() bool {
	return cd.narrowness == CQ_Narrow
}

// VariantAt returns the variant at the given index.
//
// Parameters:
//   - index: the index of the variant to return.
//
// Returns:
//   - rune: the variant at the given index.
//   - bool: true if the variant exists, false otherwise.
func (cd CharacterDescription) VariantAt(index int) (rune, bool) {
	if index < 0 || index >= len(cd.variants) {
		return 0, false
	}

	return cd.variants[index], true
}
