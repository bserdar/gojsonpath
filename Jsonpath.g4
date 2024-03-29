grammar Jsonpath;


path
  : (root_selector | current_element_selector) ( path_element )*
  ;

root_selector
  : '$'
  ;

current_element_selector
  : '@'
  ;

child_selector
  : '.'
  ;

recursive_descent
  : '..'
  ;

path_element
  : bracketed_selector
  | child_selector bracketed_selector
  | child_selector selector
  | recursive_descent bracketed_selector
  | recursive_descent selector
  | recursive_descent
  ;

bracketed_selector
  : '[' (SP)?  union (SP)? ']'
  ;


union
  : unionPart (',' unionPart)*
  ;

unionPart
  : selector
  | slice
  ;

slice
  : (singleExpression)? ':' (singleExpression)? (( ':' )? (singleExpression)?)
  ;
  
selector
  : singleExpression
  | '*'
  ;

arguments
    : '(' (argument (',' argument)* ','?)? ')'
    ;

argument
    : Ellipsis? (singleExpression | identifier)
    ;

expressionSequence
    : singleExpression (',' singleExpression)*
    ;



singleExpression
    : singleExpression (SP)? '.' (SP)? singleExpression                # ChainExpression
    | singleExpression '[' expressionSequence ']'                      # MemberIndexExpression
    | singleExpression arguments                                       # ArgumentsExpression

    | '+'  (SP)? singleExpression                                                 # UnaryPlusExpression
    | '-'  (SP)? singleExpression                                                 # UnaryMinusExpression
    | '~'  (SP)? singleExpression                                                 # BitNotExpression
    | '!'  (SP)? singleExpression                                                 # NotExpression
    | <assoc = right> singleExpression  (SP)? '**'  (SP)? singleExpression               # PowerExpression
    | singleExpression (SP)? ('*' | '/' | '%') (SP)? singleExpression                  # MultiplicativeExpression
    | singleExpression (SP)? ('+' | '-') (SP)? singleExpression                        # AdditiveExpression
    | singleExpression (SP)? '??' (SP)? singleExpression                               # CoalesceExpression
    | singleExpression (SP)? ('<' | '>' | '<=' | '>=') (SP)? singleExpression          # RelationalExpression
    | singleExpression  (SP)? In  (SP)? singleExpression                                 # InExpression
    | singleExpression  (SP)? ('==' | '!=' | '===' | '!==')  (SP)? singleExpression      # EqualityExpression
    | singleExpression  (SP)? '&' (SP)?  singleExpression                                # BitAndExpression
    | singleExpression  (SP)? '^'  (SP)? singleExpression                                # BitXOrExpression
    | singleExpression  (SP)? '|'  (SP)? singleExpression                                # BitOrExpression
    | singleExpression (SP)? '&&' (SP)? singleExpression                               # LogicalAndExpression
    | singleExpression (SP)? '||' (SP)? singleExpression                               # LogicalOrExpression
    | singleExpression  (SP)? '?' singleExpression  (SP)? ':'  (SP)? singleExpression           # TernaryExpression
    | path                                                                 # PathExpression
    | identifier                                                           # IdentifierExpression
    | literal                                                              # LiteralExpression
    | arrayLiteral                                                         # ArrayLiteralExpression
    | objectLiteral                                                        # ObjectLiteralExpression


    | '(' singleExpression ')'                                             # ParenthesizedExpression
    | '?(' singleExpression ')'                                            # FilterExpression
    ;

arrayLiteral
    : ('['  (SP)? elementList  (SP)? ']')
    ;

elementList
    : ','*  (SP)? arrayElement? ( (SP)? ','+  (SP)? arrayElement) * ','* // Yes, everything is optional
    ;

arrayElement
    : Ellipsis? singleExpression
    ;

objectLiteral
    : '{' (propertyAssignment (',' propertyAssignment)* ','?)? '}'
    ;

propertyAssignment
    : propertyName ':' singleExpression                                  # PropertyExpressionAssignment
    | '[' singleExpression ']' ':' singleExpression                      # ComputedPropertyExpressionAssignment
    | Ellipsis? singleExpression                                         # PropertyShorthand
    ;

propertyName
    : identifier
    | StringLiteral
    | numericLiteral
    | '[' singleExpression ']'
    ;


literal
    : NullLiteral
    | BooleanLiteral
    | StringLiteral
    | numericLiteral
    ;


identifier
    : Identifier
    | '@'
    ;

Identifier: IdentifierStart IdentifierPart*;


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
    | '.' [0-9] [0-9_]* ExponentPart?
    | DecimalIntegerLiteral ExponentPart?
;

StringLiteral:
    ('"' DoubleStringCharacter* '"' | '\'' SingleStringCharacter* '\'') 
;

HexIntegerLiteral    : '0' [xX] [0-9a-fA-F] HexDigit*;
OctalIntegerLiteral2 : '0' [oO] [0-7] [_0-7]*;
BinaryIntegerLiteral : '0' [bB] [01] [_01]*;

In         : 'in';

SP: [ \n\t\r]+;


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
fragment IdentifierPart: IdentifierStart | [\p{Mn}] | [\p{Nd}] | [\p{Pc}] | '\u200C' | '\u200D';
