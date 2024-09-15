package pkg

type SyllableBuilder struct {
	elems []*Character
}

// Append adds a character to the syllable.
//
// Parameters:
//   - char: the character to add to the syllable.
//
// Does nothing if the receiver or the character are nil.
func (sb *SyllableBuilder) Append(char *Character) {
	if sb == nil || char == nil {
		return
	}

	sb.elems = append(sb.elems, char)
}

// Size returns the size of the syllable.
//
// Returns:
//   - int: the size of the syllable.
func (sb SyllableBuilder) Size() int {
	return len(sb.elems)
}

/*
	private function modifierMatch matches to check if a modifier can be added to the current syllable.

Parameters:

	syllable: the current syllable
	prev: the previous character

Returns:

	int: 0 if the modifier can be added to the current syllable, 1 if the modifier can't be added to the current syllable and must be added as a new syllable, -1 otherwise
*/
func (sb SyllableBuilder) modifierMatch(prev *Character) bool {
	// A modifier can't be the first character of a syllable
	// A modifier can only be second or third character of a syllable (after a consonant)
	if prev == nil || (len(sb.elems) == 0 || len(sb.elems) > 2) {
		return false
	}

	// A modifier can only follow a vowel
	if prev.GetType() != CT_Vowel {
		return false
	}

	return true
}

/*
	private function helperMatch matches to check if a helper character can be added to the current syllable.

Parameters:

	syllable: the current syllable
	prev: the previous character

Returns:

	int: 0 if the helper character can be added to the current syllable, 1 if the helper character can't be added to the current syllable and must be added as a new syllable, -1 otherwise
*/
func (sb SyllableBuilder) helperMatch(prev *Character) bool {
	rom := sb.elems[0].GetRomanization()

	if rom != "|" && rom != "." {
		return false
	}

	// TODO: Fix this

	return true
}

func (sb SyllableBuilder) Build() *Syllable {
	if len(sb.elems) == 0 {
		return &Syllable{}
	}

	elems := make([]*Character, len(sb.elems))
	copy(elems, sb.elems)

	return &Syllable{
		chars: elems,
	}
}

func (sb *SyllableBuilder) Reset() {
	if sb == nil {
		return
	}

	if len(sb.elems) > 0 {
		for i := 0; i < len(sb.elems); i++ {
			sb.elems[i] = nil
		}

		sb.elems = sb.elems[:0]
	}
}
