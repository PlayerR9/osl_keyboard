// Code generated by "stringer -type=TokenType"; DO NOT EDIT.

package parsing

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EttEOF-0]
	_ = x[TttVowel-1]
	_ = x[TttOnset-2]
	_ = x[TttCoda-3]
	_ = x[TttDot-4]
	_ = x[TttTone-5]
	_ = x[TttPipe-6]
	_ = x[TttNewline-7]
	_ = x[NttSource-8]
	_ = x[NttSource1-9]
	_ = x[NttParagraph-10]
	_ = x[NttSentence-11]
	_ = x[NttSentence1-12]
	_ = x[NttWords-13]
	_ = x[NttWord-14]
	_ = x[NttSyllable-15]
}

const _TokenType_name = "EttEOFTttVowelTttOnsetTttCodaTttDotTttToneTttPipeTttNewlineNttSourceNttSource1NttParagraphNttSentenceNttSentence1NttWordsNttWordNttSyllable"

var _TokenType_index = [...]uint8{0, 6, 14, 22, 29, 35, 42, 49, 59, 68, 78, 90, 101, 113, 121, 128, 139}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
