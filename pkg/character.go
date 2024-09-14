package pkg

// TO DO: add space as a meaningful character

type Character struct {
	description *CharacterDescription
	varIndex    int
}

func (c *Character) String() string {
	variant := c.description.GetVariant(c.varIndex)

	return string(variant)
}

func NewCharacter(description *CharacterDescription, varIndex int) *Character {
	return &Character{
		description: description,
		varIndex:    varIndex,
	}
}

func (c *Character) IsUpper() bool {
	variant := c.description.GetVariant(c.varIndex)

	return variant >= '\uE000' &&
		variant <= '\uE018'
}

func (c *Character) GetRomanization() string {
	return c.description.GetRomanization()
}

func (c *Character) GetType() CharacterType {
	return c.description.GetType()
}

func (c *Character) IsNarrow() bool {
	return c.description.IsNarrow()
}

func (c *Character) SetVariant(varIndex int) {
	c.varIndex = varIndex
}
