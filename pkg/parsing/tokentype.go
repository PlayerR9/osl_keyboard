package parsing

//go:generate stringer -type=TokenType

type TokenType int

const (
	EttEOF TokenType = iota

	TttVowel
	TttOnset
	TttCoda

	TttDot
	TttTone
	TttPipe
	TttNewline

	NttSource
	NttSource1
	NttParagraph
	NttSentence
	NttSentence1
	NttWords
	NttWord
	NttSyllable
)
