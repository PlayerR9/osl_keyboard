parser grammar Parser;

options {
   tokenVocab = Lexer;
};

source : paragraph (NEWLINE paragraph)* EOF ;

paragraph : sentence (DOT sentence)* ;

sentence : TONE? word+ (PIPE word+)* ;

word : syllable (syllable syllable?)? ;

syllable
   : ONSET? VOWEL CODA?
   | ONSET
   ;