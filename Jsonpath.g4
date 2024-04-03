grammar Jsonpath;

path: expression (SP)? EOF ;

simplePath: '/' simplePathExpr (SP)? EOF;

simplePathExpr
  : simplePathExpr (SP)? '/' (SP)? simplePathExpr
  | literal                    
  | identifier
  | '*'
  ;
  
selector
  : '$'
  | '@'
  | '*'
  ;

arguments
    : '(' (expression (',' expression)* ','?)? ')'
    ;

slice
   : (expression)? ':' (expression)? (( ':' )? (expression)?)
   ;

indexExpression
   : slice
   | expression
   ;


expression
    : expression (SP)? '.' (SP)? expression                                                             # ChainExpression
    | expression (SP)? '..' (SP)? expression                                                            # RecursiveDescentExpression
    | expression (SP)? '..'                                                                             # RecursiveDescentTermExpression
    | expression (SP)? '[' (SP)? indexExpression ( (SP)? ',' (SP)? indexExpression)* (SP)? ']'          # MemberIndexExpression
    | expression (SP)? '..' (SP)? '[' (SP)? indexExpression ( (SP)? ',' (SP)? indexExpression)* (SP)? ']'          # RecursiveDescentMemberIndexExpression

    | expression arguments                                                  # ArgumentsExpression

    | '+'  (SP)? expression                                                 # UnaryPlusExpression
    | '-'  (SP)? expression                                                 # UnaryMinusExpression
    | '~'  (SP)? expression                                                 # BitNotExpression
    | '!'  (SP)? expression                                                 # NotExpression
    
    | <assoc = right> expression  (SP)? '**'  (SP)? expression               # PowerExpression
    | expression (SP)? ('*' | '/' | '%') (SP)? expression                    # MultiplicativeExpression
    | expression (SP)? ('+' | '-') (SP)? expression                          # AdditiveExpression
    | expression (SP)? ('<' | '>' | '<=' | '>=') (SP)? expression            # RelationalExpression
    | expression  (SP)? 'in'  (SP)? expression                                 # InExpression
    | expression  (SP)? ('==' | '!=' | '===' | '!==')  (SP)? expression      # EqualityExpression
    | expression  (SP)? '&' (SP)?  expression                                # BitAndExpression
    | expression  (SP)? '^'  (SP)? expression                                # BitXOrExpression
    | expression  (SP)? '|'  (SP)? expression                                # BitOrExpression
    | expression (SP)? '&&' (SP)? expression                                 # LogicalAndExpression
    | expression (SP)? '||' (SP)? expression                                 # LogicalOrExpression

    | literal                                                              # LiteralExpression
    | arrayLiteral                                                         # ArrayLiteralExpression
    | identifier                                                           # IdentifierExpression
    | selector                                                             # SelectorExpression


    | '(' expression ')'                                             # ParenthesizedExpression
    | '?(' expression ')'                                            # FilterExpression
    ;

arrayLiteral
    : ('['  (SP)? elementList  (SP)? ']')
    ;

elementList
    : ','*  (SP)? expression? ( (SP)? ','+  (SP)? expression) * ','* // Yes, everything is optional
    ;


literal
    : NullLiteral
    | BooleanLiteral
    | numericLiteral
    | StringLiteral
    ;


identifier
    : Identifier
    ;



numericLiteral
    : DecimalLiteral
    | HexIntegerLiteral
    | OctalIntegerLiteral2
    | BinaryIntegerLiteral
    ;


Ellipsis                   : '...';
NullLiteral: 'null';
BooleanLiteral: 'true' | 'false';
DecimalLiteral:
    DecimalIntegerLiteral '.' [0-9] [0-9_]* ExponentPart?
    | DecimalIntegerLiteral ExponentPart?
;

StringLiteral:
    ('"' DoubleStringCharacter* '"' | '\'' SingleStringCharacter* '\'') 
;

HexIntegerLiteral    : '0' [xX] [0-9a-fA-F] HexDigit*;
OctalIntegerLiteral2 : '0' [oO] [0-7] [_0-7]*;
BinaryIntegerLiteral : '0' [bB] [01] [_01]*;


SP: [ \n\t\r]+;

Identifier: IdentifierStart IdentifierPart*;


fragment ExponentPart: [eE] [+-]? [0-9_]+;
fragment DecimalIntegerLiteral: '0' | [0-9] [0-9_]*;
fragment DoubleStringCharacter: ~["\\\r\n] | '\\' EscapeSequence | LineContinuation;
fragment SingleStringCharacter: ~['\\\r\n] | '\\' EscapeSequence | LineContinuation;
fragment EscapeSequence:
    CharacterEscapeSequence
    | '0' // no digit ahead! TODO
    | HexEscapeSequence
    | UnicodeEscapeSequence
    | ExtendedUnicodeEscapeSequence
;

fragment CharacterEscapeSequence: SingleEscapeCharacter | NonEscapeCharacter;

fragment HexEscapeSequence: 'x' HexDigit HexDigit;

fragment UnicodeEscapeSequence:
    'u' HexDigit HexDigit HexDigit HexDigit
    | 'u' '{' HexDigit HexDigit+ '}'
;

fragment ExtendedUnicodeEscapeSequence: 'u' '{' HexDigit+ '}';

fragment SingleEscapeCharacter: ['"\\bfnrtv];
fragment HexDigit: [_0-9a-fA-F];
fragment NonEscapeCharacter: ~['"\\bfnrtv0-9xu\r\n];
fragment LineContinuation: '\\' [\r\n\u2028\u2029]+;
fragment IdentifierStart: [\p{L}] | [$_] | '\\' UnicodeEscapeSequence;
fragment IdentifierPart: IdentifierStart | '-' | [\p{Mn}] | [\p{Nd}] | [\p{Pc}] | '\u200C' | '\u200D';
