package pkg

import "unicode/utf8"

type CharacterType int8

const (
	// Character types
	CT_Vowel CharacterType = iota
	CT_Consonant
	CT_Coda
	CT_Extra
	CT_Helper
)

type CharacterNarrowness int8

const (
	// Character qualities
	CQ_Narrow CharacterNarrowness = iota
	CQ_Wide
)

type CharacterDescription struct {
	romanization string
	len          int
	ctype        CharacterType // character type
	isNarrow     CharacterNarrowness
	variants     []rune
}

func NewCharacterDescription(romanization string, ctype CharacterType, isNarrow CharacterNarrowness, variants ...rune) *CharacterDescription {
	return &CharacterDescription{
		romanization: romanization,
		len:          utf8.RuneCountInString(romanization),
		ctype:        ctype,
		isNarrow:     isNarrow,
		variants:     variants,
	}
}

func (d *CharacterDescription) Size() int {
	return d.len
}

func (d *CharacterDescription) GetRomanization() string {
	return d.romanization
}

func (d *CharacterDescription) GetType() CharacterType {
	return d.ctype
}

func (d *CharacterDescription) VariantSize() int {
	return len(d.variants)
}

func (d *CharacterDescription) IsNarrow() bool {
	return d.isNarrow == CQ_Narrow
}

func (d *CharacterDescription) GetVariant(index int) rune {
	return d.variants[index]
}
