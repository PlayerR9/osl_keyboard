lexer grammar Lexer;

VOWEL : [uöyiëä] | [oea][\u0308]? ;
ONSET : [rsvmkbntd] ;
CODA : [ṇṃl] | [nm][\u0323] ;

DOT : '.' ;

TONE : '\\' [mhrl] ;

PIPE : '|' ;

NEWLINE : ('\r'? '\n')+ ;
WS : [ \t]+ -> skip ;
 