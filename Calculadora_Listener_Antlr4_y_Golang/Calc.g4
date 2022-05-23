 grammar Calc;
// antlr4 -Dlanguage=Go -o parser Calc.g4
//tokens
MUL: '*';
ADD: '+';
LB: '(';
RB: ')';
DIGIT: [0-9]+;
WS: [ \r\n\t]+ -> skip;

//rules
l   : e EOF;

e   : e '+' t   # Sum
    | t         # PassT
    ;

t   : t '*' f   # Mul
    | f         # PassF
    ;


f   : '(' e ')' # PassE
    | DIGIT     # Digit
    ;
