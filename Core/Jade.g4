grammar Jade;

/* Jade Lexer */
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LLIST: '[';
RLIST: ']';
ASSIGN: '=';
DOT: '.';
COMMA: ',';
COLON: ':';
SEMI: ';';
LARROW: '->';
RARROW: '<-';
ARROWASSIGN: '=>';
QUESTION: '?';

// Predecence Operators
NOT: '!' | 'not';
PREDTWO: '+' | '-' | '**';
PREDONE: '*' | '/' | '%';
COMPARATIVE: '==' | '!=' | '~=' | '>' | '<' | '<=' | '>=' | 'and' | '&&' | 'or' | '||';

/// Keywords
FUNC: 'func';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
IMPORT: 'import';
FROM: 'from';
OVERRIDE: 'override';
NEW: 'new';
CLASS: 'class';
PUBLIC: 'public';
PRIVATE: 'private';
RETURN: 'return';
BREAK: 'break';
CONTINUE: 'continue';
UNDEFINE: 'undefine';
CONST: 'const';

/// Data type Keywords
ALL_KWS: 'string' | 'int' | 'float' | 'bool';

/// Skipped Rules
WS: [ \t\r\n]+ -> skip;
COMMENT: ('//' | '<--' | '#') ~[\r\n]* -> skip;
MULTILINECOMMENT: '/*' .*? '*/' -> skip;

/// Data Types
BOOL: ('true' | 'false');
NUL: 'null';

APOSTROPHE: '\'';

ID: [a-zA-Z_] [a-zA-Z_0-9]*;
INT: '-'? [0-9]+;
FLOAT: '-'? [0-9]* '.' [0-9]+;
STRING: ('"' | APOSTROPHE) .*? ('"' | APOSTROPHE);

/* Jade Parser */
parse: stmt* EOF;

block: LBRACE stmt* RBRACE | stmt;

stmt
    : call
    | varAssign
    | funcAssign
    | expr
    | classdef
    | allStmts
    | iterationStmt
    | functionStmt
    ;

iterationStmt: BREAK | CONTINUE;
functionStmt: RETURN expr;

allStmts
    : ifStmt
    | whileStmt
    | undefineStmt
    | importStmt
    ;

importStmt
    : IMPORT STRING (COMMA STRING)* FROM STRING
    ;

ifStmt
    : IF condition block (ELSE IF condition block)* (ELSE block)?
    ;

whileStmt
    : WHILE condition block
    ;

condition
    : LPAREN expr RPAREN
    | expr
    ;

undefineStmt: UNDEFINE LPAREN ID RPAREN;

inheritList: ID (COMMA ID)*;

classdef: CLASS ID RARROW inheritList?;

args: expr (COMMA expr)*;
params: ALL_KWS? ID (COMMA ALL_KWS? ID)*;

call: ID LPAREN args? RPAREN;

varAssign: (PUBLIC | PRIVATE)? CONST? ALL_KWS? ID ASSIGN expr;
funcAssign
    : (PUBLIC | PRIVATE)? (ALL_KWS? | FUNC) OVERRIDE? ID LPAREN params? RPAREN block
    | (PUBLIC | PRIVATE)? ALL_KWS? ID LPAREN params? RPAREN ARROWASSIGN block
    ;

getAttributes: atom (DOT ID LPAREN args? RPAREN)*;

funcExpr: ID LPAREN params? RPAREN ARROWASSIGN block;

expr
    : atom
    | call
    | op=NOT expr
    | expr op=PREDONE expr
    | expr op=PREDTWO expr
    | expr op=COMPARATIVE expr
    | getAttributes
    | funcExpr
    | LPAREN expr RPAREN
    ;

array: LBRACE args? RBRACE;

atom: array | ID | INT | FLOAT | STRING | NUL | BOOL;
