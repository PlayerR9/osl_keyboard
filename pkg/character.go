package pkg

import (
	gcers "github.com/PlayerR9/go-commons/errors"
	gcint "github.com/PlayerR9/go-commons/ints"
	dba "github.com/PlayerR9/go-debug/assert"
)

// Character is a collection of characters.
type Character struct {
	// description is the description of the character.
	description *CharacterDescription

	// var_idx is the index of the variant of the character.
	var_idx int
}

// String returns the romanized character.
//
// Returns:
//   - string: the romanized character.
func (c Character) String() string {
	variant, ok := c.description.VariantAt(c.var_idx)
	dba.AssertOk(ok, "c.description.GetVariant(%d)", c.var_idx)

	return string(variant)
}

// NewCharacter creates a new character.
//
// Parameters:
//   - description: the description of the character.
//   - varIndex: the index of the variant of the character.
//
// Returns:
//   - *Character: the new character.
//   - error: an error if the description is nil or if the index is out of bounds.
func NewCharacter(description *CharacterDescription, var_idx int) (*Character, error) {
	if description == nil {
		return nil, gcers.NewErrNilParameter("description")
	}

	size := description.VariantSize()

	if var_idx < 0 || var_idx >= size {
		return nil, gcers.NewErrInvalidParameter("var_idx", gcint.NewErrOutOfBounds(var_idx, 0, size))
	}

	return &Character{
		description: description,
		var_idx:     var_idx,
	}, nil
}

// IsUpper checks if the character is upper.
//
// Returns:
//   - bool: true if the character is upper, false otherwise.
func (c Character) IsUpper() bool {
	variant, ok := c.description.VariantAt(c.var_idx)
	dba.AssertOk(ok, "c.description.GetVariant(%d)", c.var_idx)

	return variant >= '\uE000' &&
		variant <= '\uE018'
}

// GetRomanization returns the romanization of the character.
//
// Returns:
//   - string: the romanization of the character.
func (c Character) GetRomanization() string {
	return c.description.Romanization
}

// GetType returns the type of the character.
//
// Returns:
//   - CharacterType: the type of the character.
func (c *Character) GetType() CharacterType {
	return c.description.Type
}

// IsNarrow checks if the character is narrow.
//
// Returns:
//   - bool: true if the character is narrow, false otherwise.
func (c Character) IsNarrow() bool {
	return c.description.IsNarrow()
}

// SetVariant sets the variant of the character.
//
// Parameters:
//   - var_idx: the index of the variant of the character.
//
// Returns:
//   - error: an error if the index is out of bounds.
//
// Does nothing if the receiver is nil.
func (c *Character) SetVariant(var_idx int) error {
	if c == nil {
		return nil
	}

	if var_idx < 0 || var_idx >= c.description.VariantSize() {
		return gcers.NewErrInvalidParameter("varIndex", gcint.NewErrOutOfBounds(var_idx, 0, c.description.VariantSize()))
	}

	c.var_idx = var_idx

	return nil
}
