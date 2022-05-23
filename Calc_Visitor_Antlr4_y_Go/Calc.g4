grammar Calc;

//comando de compilacion de antlr
//antlr4 -Dlanguage=Go -o parser -visitor -no-listener Calc.g4


//token
INT :   [0-9]+;
WS  :   [ \t]+  ->  skip;

//rules
expr    :   left=expr op=('*'|'/') right=expr   # Op 
        |   left=expr op=('+'|'-') right=expr   # Op 
        |   digit=INT                           # Digit 
        |   '(' expr ')'                        # Paren
        ;
