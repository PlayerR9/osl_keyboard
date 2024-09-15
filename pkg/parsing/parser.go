package parsing

import (
	gr "github.com/PlayerR9/SlParser/grammar"
	prx "github.com/PlayerR9/SlParser/parser"
	dba "github.com/PlayerR9/go-debug/assert"
)

var (
	// source : source1 EOF ;
	rule1 *prx.Rule[TokenType]

	// source1 : paragraph ;
	rule2 *prx.Rule[TokenType]

	// source1 : paragraph NEWLINE source1 ;
	rule3 *prx.Rule[TokenType]

	// sentence : sentence1 ;
	rule4 *prx.Rule[TokenType]

	// sentence : TONE sentence1 ;
	rule5 *prx.Rule[TokenType]

	// sentence1 : words ;
	rule6 *prx.Rule[TokenType]

	// sentence1 : words PIPE sentence1 ;
	rule7 *prx.Rule[TokenType]

	// words : word ;
	rule8 *prx.Rule[TokenType]

	// words : word words ;
	rule9 *prx.Rule[TokenType]

	// word : syllable ;
	rule10 *prx.Rule[TokenType]

	// word : syllable syllable ;
	rule11 *prx.Rule[TokenType]

	// word : syllable syllable syllable ;
	rule12 *prx.Rule[TokenType]

	// syllable : VOWEL ;
	rule13 *prx.Rule[TokenType]

	// syllable : VOWEL CODA ;
	rule14 *prx.Rule[TokenType]

	// syllable : ONSET VOWEL ;
	rule15 *prx.Rule[TokenType]

	// syllable : ONSET VOWEL CODA ;
	rule16 *prx.Rule[TokenType]

	// syllable : ONSET ;
	rule17 *prx.Rule[TokenType]

	// paragraph : sentence ;
	rule18 *prx.Rule[TokenType]

	// paragraph : sentence DOT paragraph ;
	rule19 *prx.Rule[TokenType]
)

func init() {
	var err error

	rule1, err = prx.NewRule(NttSource, NttSource1, EttEOF)
	dba.AssertErr(err, "parser.NewRule(NttSource, NttSource1, EttEOF)")

	// source1 : paragraph ;
	rule2, err = prx.NewRule(NttSource1, NttSentence)
	dba.AssertErr(err, "parser.NewRule(NttSource1, NttSentence)")

	// source1 : paragraph NEWLINE paragraph ;
	rule3, err = prx.NewRule(NttSource1, NttSentence, TttNewline, NttSource1)
	dba.AssertErr(err, "parser.NewRule(NttSource1, NttSentence, TttNewline, NttSource1)")

	// sentence : sentence1 ;
	rule4, err = prx.NewRule(NttSentence, NttSentence1)
	dba.AssertErr(err, "parser.NewRule(NttSentence, NttSentence1)")

	// sentence : TONE sentence1 ;
	rule5, err = prx.NewRule(NttSentence, TttTone, NttSentence1)
	dba.AssertErr(err, "parser.NewRule(NttSentence, TttTone, NttSentence1)")

	// sentence1 : words ;
	rule6, err = prx.NewRule(NttSentence1, NttWords)
	dba.AssertErr(err, "parser.NewRule(NttSentence1, NttWords)")

	// sentence1 : words PIPE sentence1 ;
	rule7, err = prx.NewRule(NttSentence1, NttWords, TttPipe, NttSentence1)
	dba.AssertErr(err, "parser.NewRule(NttSentence1, NttWords, TttPipe, NttSentence1)")

	// words : word ;
	rule8, err = prx.NewRule(NttWords, NttWord)
	dba.AssertErr(err, "parser.NewRule(NttWords, NttWord)")

	// words : word words ;
	rule9, err = prx.NewRule(NttWords, NttWord, NttWords)
	dba.AssertErr(err, "parser.NewRule(NttWords, NttWord, NttWords)")

	// word : syllable ;
	rule10, err = prx.NewRule(NttWord, NttSyllable)
	dba.AssertErr(err, "parser.NewRule(NttWord, NttSyllable)")

	// word : syllable syllable ;
	rule11, err = prx.NewRule(NttWord, NttSyllable, NttSyllable)
	dba.AssertErr(err, "parser.NewRule(NttWord, NttSyllable, NttSyllable)")

	// word : syllable syllable syllable ;
	rule12, err = prx.NewRule(NttWord, NttSyllable, NttSyllable, NttSyllable)
	dba.AssertErr(err, "parser.NewRule(NttWord, NttSyllable, NttSyllable, NttSyllable)")

	// syllable : VOWEL ;
	rule13, err = prx.NewRule(NttSyllable, TttVowel)
	dba.AssertErr(err, "parser.NewRule(NttSyllable, TttVowel)")

	// syllable : VOWEL CODA ;
	rule14, err = prx.NewRule(NttSyllable, TttVowel, TttCoda)
	dba.AssertErr(err, "parser.NewRule(NttSyllable, TttVowel, TttCoda)")

	// syllable : ONSET VOWEL ;
	rule15, err = prx.NewRule(NttSyllable, TttOnset, TttVowel)
	dba.AssertErr(err, "parser.NewRule(NttSyllable, TttOnset, TttVowel)")

	// syllable : ONSET VOWEL CODA ;
	rule16, err = prx.NewRule(NttSyllable, TttOnset, TttVowel, TttCoda)
	dba.AssertErr(err, "parser.NewRule(NttSyllable, TttOnset, TttVowel, TttCoda)")

	// syllable : ONSET ;
	rule17, err = prx.NewRule(NttSyllable, TttOnset)
	dba.AssertErr(err, "parser.NewRule(NttSyllable, TttOnset)")

	// paragraph : sentence ;
	rule18, err = prx.NewRule(NttParagraph, NttSentence)
	dba.AssertErr(err, "parser.NewRule(NttParagraph, NttSentence)")

	// paragraph : sentence DOT paragraph ;
	rule19, err = prx.NewRule(NttParagraph, NttSentence, TttDot, NttParagraph)
	dba.AssertErr(err, "parser.NewRule(NttParagraph, NttSentence, TttDot, NttParagraph)")
}

var (
	parser prx.Parser[TokenType]
)

func init() {
	builder := prx.NewBuilder[TokenType]()

	// source1 : paragraph NEWLINE source1 # ; (3)
	// source : source1 # EOF ; (1)
	builder.Register(NttSource1, prx.UnambiguousRule(
		prx.MustNewItem(rule3, 3),
		prx.MustNewItem(rule1, 1),
	))

	// source : source1 EOF # ; (1)
	builder.Register(EttEOF, prx.UnambiguousRule(prx.MustNewItem(rule1, 2)))

	builder.Register(NttParagraph, func(parser *prx.Parser[TokenType], top1, lookahead *gr.Token[TokenType]) ([]*prx.Item[TokenType], error) {
		has_la := prx.CheckLookahead(lookahead, TttNewline)

		var it1 *prx.Item[TokenType]
		var err error

		_, has_top2 := prx.CheckTop(parser, TttDot)
		if has_top2 {
			// paragraph : sentence DOT paragraph # ; (19)

			it1, err = prx.NewItem(rule19, 3)
			dba.AssertErr(err, "parser.NewItem(rule19, %d)", 3)
		} else if has_la {
			// source1 : paragraph # NEWLINE source1 ; (3)

			it1, err = prx.NewItem(rule3, 1)
			dba.AssertErr(err, "parser.NewItem(rule3, %d)", 1)
		} else {
			// source1 : paragraph # ; (2)

			it1, err = prx.NewItem(rule2, 1)
			dba.AssertErr(err, "parser.NewItem(rule2, %d)", 1)
		}

		return []*prx.Item[TokenType]{it1}, nil
	})

	// source1 : paragraph NEWLINE # source1 ; (3)
	builder.Register(TttNewline, prx.UnambiguousRule(prx.MustNewItem(rule3, 2)))

	// sentence : sentence1 # ; (4)
	// sentence : TONE sentence1 # ; (5)
	// sentence1 : words PIPE sentence1 # ; (7)
	builder.Register(NttSentence1, prx.UnambiguousRule(
		prx.MustNewItem(rule4, 1),
		prx.MustNewItem(rule5, 2),
		prx.MustNewItem(rule7, 3),
	))

	// sentence : TONE # sentence1 ; (5)
	builder.Register(TttTone, prx.UnambiguousRule(prx.MustNewItem(rule5, 1)))

	builder.Register(NttWords, func(parser *prx.Parser[TokenType], top1, lookahead *gr.Token[TokenType]) ([]*prx.Item[TokenType], error) {
		has_la := prx.CheckLookahead(lookahead, TttPipe)

		var it1 *prx.Item[TokenType]
		var err error

		_, has_top2 := prx.CheckTop(parser, NttWord)
		if has_top2 {
			// words : word words # ; (9)

			it1, err = prx.NewItem(rule9, 2)
			dba.AssertErr(err, "parser.NewItem(rule9, %d)", 2)
		} else if has_la {
			// sentence1 : words # PIPE sentence1 ; (7)

			it1, err = prx.NewItem(rule7, 1)
			dba.AssertErr(err, "parser.NewItem(rule7, %d)", 1)
		} else {
			// sentence1 : words # ; (6)

			it1, err = prx.NewItem(rule6, 1)
			dba.AssertErr(err, "parser.NewItem(rule6, %d)", 1)
		}

		return []*prx.Item[TokenType]{it1}, nil
	})

	// sentence1 : words PIPE # sentence1 ; (7)
	builder.Register(TttPipe, prx.UnambiguousRule(prx.MustNewItem(rule7, 2)))

	builder.Register(NttWord, func(parser *prx.Parser[TokenType], top1, lookahead *gr.Token[TokenType]) ([]*prx.Item[TokenType], error) {
		var it1 *prx.Item[TokenType]
		var err error

		has_la := prx.CheckLookahead(lookahead, TttVowel, TttOnset)
		if has_la {
			// words : word # words ; (9)
			// -- VOWEL
			// -- ONSET

			it1, err = prx.NewItem(rule9, 1)
			dba.AssertErr(err, "parser.NewItem(rule9, %d)", 1)
		} else {
			// words : word # ; (8)

			it1, err = prx.NewItem(rule8, 1)
			dba.AssertErr(err, "parser.NewItem(rule8, %d)", 1)
		}

		return []*prx.Item[TokenType]{it1}, nil
	})

	builder.Register(NttSyllable, func(parser *prx.Parser[TokenType], top1, lookahead *gr.Token[TokenType]) ([]*prx.Item[TokenType], error) {
		has_la := prx.CheckLookahead(lookahead, TttVowel, TttOnset)

		var it1 *prx.Item[TokenType]
		var err error

		_, has_top2 := prx.CheckTop(parser, NttSyllable)
		if has_top2 {
			_, has_top3 := prx.CheckTop(parser, NttSyllable)
			if has_top3 {
				// word : syllable syllable syllable # ; (12)

				it1, err = prx.NewItem(rule12, 3)
				dba.AssertErr(err, "parser.NewItem(rule12, %d)", 3)
			} else if has_la {
				// word : syllable syllable # syllable ; (12)

				it1, err = prx.NewItem(rule12, 2)
				dba.AssertErr(err, "parser.NewItem(rule12, %d)", 2)
			} else {
				// word : syllable syllable # ; (11)

				it1, err = prx.NewItem(rule11, 2)
				dba.AssertErr(err, "parser.NewItem(rule11, %d)", 2)
			}
		} else if has_la {
			has_la2 := prx.CheckLookahead(lookahead.Lookahead, TttVowel, TttOnset)
			if has_la2 {
				// word : syllable # syllable syllable ; (12)

				it1, err = prx.NewItem(rule12, 1)
				dba.AssertErr(err, "parser.NewItem(rule12, %d)", 1)
			} else {
				// word : syllable # syllable ; (11)

				it1, err = prx.NewItem(rule11, 1)
				dba.AssertErr(err, "parser.NewItem(rule11, %d)", 1)
			}
		} else {
			// word : syllable # ; (10)

			it1, err = prx.NewItem(rule10, 1)
			dba.AssertErr(err, "parser.NewItem(rule10, %d)", 1)
		}

		return []*prx.Item[TokenType]{it1}, nil
	})

	builder.Register(TttVowel, func(parser *prx.Parser[TokenType], top1, lookahead *gr.Token[TokenType]) ([]*prx.Item[TokenType], error) {
		var it1 *prx.Item[TokenType]
		var err error

		has_la := prx.CheckLookahead(lookahead, TttCoda)

		_, has_top2 := prx.CheckTop(parser, TttOnset)
		if has_top2 {
			if has_la {
				// syllable : ONSET VOWEL # CODA ; (16)

				it1, err = prx.NewItem(rule16, 2)
				dba.AssertErr(err, "parser.NewItem(rule16, %d)", 2)
			} else {
				// syllable : ONSET VOWEL # ; (15)

				it1, err = prx.NewItem(rule15, 2)
				dba.AssertErr(err, "parser.NewItem(rule15, %d)", 2)
			}
		} else if has_la {
			// syllable : VOWEL # CODA ; (14)

			it1, err = prx.NewItem(rule14, 1)
			dba.AssertErr(err, "parser.NewItem(rule14, %d)", 1)
		} else {
			// syllable : VOWEL # ; (13)

			it1, err = prx.NewItem(rule13, 1)
			dba.AssertErr(err, "parser.NewItem(rule13, %d)", 1)
		}

		return []*prx.Item[TokenType]{it1}, nil
	})

	// syllable : ONSET VOWEL CODA # ; (16)
	// syllable : VOWEL CODA # ; (14)
	builder.Register(TttCoda, prx.UnambiguousRule(
		prx.MustNewItem(rule16, 3),
		prx.MustNewItem(rule14, 2),
	))

	builder.Register(TttOnset, func(parser *prx.Parser[TokenType], top1, lookahead *gr.Token[TokenType]) ([]*prx.Item[TokenType], error) {
		var it1 *prx.Item[TokenType]
		var err error

		has_la1 := prx.CheckLookahead(lookahead, TttVowel)
		if has_la1 {
			has_la2 := prx.CheckLookahead(lookahead, TttCoda)
			if has_la2 {
				// syllable : ONSET # VOWEL CODA ; (16)

				it1, err = prx.NewItem(rule16, 1)
				dba.AssertErr(err, "parser.NewItem(rule16, %d)", 1)
			} else {
				// syllable : ONSET # VOWEL ; (15)

				it1, err = prx.NewItem(rule15, 1)
				dba.AssertErr(err, "parser.NewItem(rule15, %d)", 1)
			}
		} else {
			// syllable : ONSET # ; (17)

			it1, err = prx.NewItem(rule17, 1)
			dba.AssertErr(err, "parser.NewItem(rule17, %d)", 1)
		}

		return []*prx.Item[TokenType]{it1}, nil
	})

	builder.Register(NttSentence, func(parser *prx.Parser[TokenType], top1, lookahead *gr.Token[TokenType]) ([]*prx.Item[TokenType], error) {
		var it1 *prx.Item[TokenType]
		var err error

		has_la := prx.CheckLookahead(lookahead, TttDot)
		if has_la {
			// paragraph : sentence # DOT paragraph ; (19)

			it1, err = prx.NewItem(rule19, 1)
			dba.AssertErr(err, "parser.NewItem(rule19, %d)", 1)
		} else {
			// paragraph : sentence # ; (18)

			it1, err = prx.NewItem(rule18, 1)
			dba.AssertErr(err, "parser.NewItem(rule18, %d)", 1)
		}

		return []*prx.Item[TokenType]{it1}, nil
	})

	// paragraph : sentence DOT # paragraph ; (19)
	builder.Register(TttDot, prx.UnambiguousRule(prx.MustNewItem(rule19, 2)))

	parser = builder.Build()
}
