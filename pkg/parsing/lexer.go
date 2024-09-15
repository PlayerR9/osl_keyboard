package parsing

import (
	"io"

	gr "github.com/PlayerR9/SlParser/grammar"
	lxr "github.com/PlayerR9/SlParser/lexer"
	dba "github.com/PlayerR9/go-debug/assert"
)

var (
	lexer lxr.Lexer[TokenType]
)

func init() {
	builder := lxr.NewBuilder[TokenType]()

	// VOWEL : [uöyiëä] | [oea][\u0308]? ;
	builder.Register('u', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttVowel, "u")
		return tk, nil
	})

	builder.Register('ö', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttVowel, "ö")
		return tk, nil
	})

	builder.Register('o', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		next, err := lexer.PeekRune()
		if err == io.EOF {
			tk := gr.NewTerminalToken(TttVowel, "o")
			return tk, nil
		} else if err == nil {
			return nil, err
		}

		var tk *gr.Token[TokenType]

		if next != '\u0308' {
			tk = gr.NewTerminalToken(TttVowel, "o")
		} else {
			_, err := lexer.NextRune()
			dba.AssertErr(err, "lexer.NextRune()")

			tk = gr.NewTerminalToken(TttVowel, "ö")
		}

		return tk, nil
	})

	builder.Register('y', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttVowel, "y")
		return tk, nil
	})

	builder.Register('ë', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttVowel, "ë")
		return tk, nil
	})

	builder.Register('e', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		next, err := lexer.PeekRune()
		if err == io.EOF {
			tk := gr.NewTerminalToken(TttVowel, "e")
			return tk, nil
		} else if err == nil {
			return nil, err
		}

		var tk *gr.Token[TokenType]

		if next != '\u0308' {
			tk = gr.NewTerminalToken(TttVowel, "e")
		} else {
			_, err := lexer.NextRune()
			dba.AssertErr(err, "lexer.NextRune()")

			tk = gr.NewTerminalToken(TttVowel, "ë")
		}

		return tk, nil
	})

	builder.Register('i', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttVowel, "i")
		return tk, nil
	})

	builder.Register('ä', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttVowel, "ä")
		return tk, nil
	})

	builder.Register('a', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		next, err := lexer.PeekRune()
		if err == io.EOF {
			tk := gr.NewTerminalToken(TttVowel, "a")
			return tk, nil
		} else if err == nil {
			return nil, err
		}

		var tk *gr.Token[TokenType]

		if next != '\u0308' {
			tk = gr.NewTerminalToken(TttVowel, "a")
		} else {
			_, err := lexer.NextRune()
			dba.AssertErr(err, "lexer.NextRune()")

			tk = gr.NewTerminalToken(TttVowel, "ä")
		}

		return tk, nil
	})

	// ONSET : [rsvmkbntd] ;
	builder.Register('r', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttOnset, "r")
		return tk, nil
	})

	builder.Register('s', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttOnset, "s")
		return tk, nil
	})

	builder.Register('v', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttOnset, "v")
		return tk, nil
	})

	builder.Register('k', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttOnset, "k")
		return tk, nil
	})

	builder.Register('b', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttOnset, "b")
		return tk, nil
	})

	builder.Register('t', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttOnset, "t")
		return tk, nil
	})

	builder.Register('d', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttOnset, "d")
		return tk, nil
	})

	// CODA : [ṇṃl] | [nm][\u0323] ;
	builder.Register('ṃ', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttCoda, "ṃ")
		return tk, nil
	})

	builder.Register('m', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		next, err := lexer.PeekRune()
		if err == io.EOF {
			tk := gr.NewTerminalToken(TttOnset, "m")
			return tk, nil
		} else if err == nil {
			return nil, err
		}

		var tk *gr.Token[TokenType]

		if next != '\u0323' {
			tk = gr.NewTerminalToken(TttOnset, "m")
		} else {
			_, err := lexer.NextRune()
			dba.AssertErr(err, "lexer.NextRune()")

			tk = gr.NewTerminalToken(TttCoda, "ṃ")
		}

		return tk, nil
	})

	builder.Register('ṇ', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttCoda, "ṇ")
		return tk, nil
	})

	builder.Register('n', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		next, err := lexer.PeekRune()
		if err == io.EOF {
			tk := gr.NewTerminalToken(TttOnset, "n")
			return tk, nil
		} else if err == nil {
			return nil, err
		}

		var tk *gr.Token[TokenType]

		if next != '\u0323' {
			tk = gr.NewTerminalToken(TttOnset, "n")
		} else {
			_, err := lexer.NextRune()
			dba.AssertErr(err, "lexer.NextRune()")

			tk = gr.NewTerminalToken(TttCoda, "ṇ")
		}

		return tk, nil
	})

	builder.Register('l', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttCoda, "l")
		return tk, nil
	})

	// DOT : '.' ;
	builder.Register('.', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttDot, ".")
		return tk, nil
	})

	// TONE : '\\' [mhrl] ;
	builder.Register('\\', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		next, err := lexer.NextRune()
		if err == io.EOF {
			return nil, lxr.NewErrUnexpectedChar('\\', []rune{'m', 'h', 'r', 'l'}, nil)
		} else if err != nil {
			return nil, err
		}

		if next != 'm' && next != 'h' && next != 'r' && next != 'l' {
			return nil, lxr.NewErrUnexpectedChar('\\', []rune{'m', 'h', 'r', 'l'}, &next)
		}

		tk := gr.NewTerminalToken(TttTone, string(next))
		return tk, nil
	})

	// PIPE : '|' ;
	builder.Register('|', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		tk := gr.NewTerminalToken(TttPipe, "|")
		return tk, nil
	})

	// NEWLINE : ('\r'? '\n')+ ;
	builder.Register('\r', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		next, err := lexer.NextRune()
		if err == io.EOF {
			return nil, lxr.NewErrUnexpectedChar('\r', []rune{'\n'}, nil)
		} else if err != nil {
			return nil, err
		}

		if next != '\n' {
			return nil, lxr.NewErrUnexpectedChar('\r', []rune{'\n'}, &next)
		}

		_, err = lxr.LexMany(lexer, lxr.FragNewline)
		if err != nil {
			return nil, err
		}

		return gr.NewTerminalToken(TttNewline, "\n"), nil
	})

	// NEWLINE : ('\r'? '\n')+ ;
	builder.Register('\n', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		_, err := lxr.LexMany(lexer, lxr.FragNewline)
		if err != nil {
			return nil, err
		}

		return gr.NewTerminalToken(TttNewline, "\n"), nil
	})

	// WS : [ \t]+ -> skip ;
	builder.Register(' ', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		_, err := lxr.LexMany(lexer, lxr.FragWs)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	builder.Register('\t', func(lexer *lxr.Lexer[TokenType], char rune) (*gr.Token[TokenType], error) {
		_, err := lxr.LexMany(lexer, lxr.FragWs)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	lexer = builder.Build()
}

/*
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
*/

// FullLex lexes the input bytes and returns the tokens.
//
// Parameters:
//   - data: the input bytes.
//
// Returns:
//   - []*gr.Token[TokenType]: the tokens extracted from the input bytes.
//   - error: an error if the input is invalid.
func FullLex(data []byte) ([]*gr.Token[TokenType], error) {
	is := lxr.NewStream().FromBytes(data)

	lexer.SetInputStream(is)

	err := lexer.Lex()
	tokens := lexer.Tokens()

	lexer.Reset()

	return tokens, err
}
