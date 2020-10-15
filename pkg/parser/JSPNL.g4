grammar JSPNL;

prog
    :	((statement)? NEWLINE)*
    ;

statement
    : declaration
    | simple
    ;

// declaration

declaration
    : node_declaration
    | arc_declaration
    | reward_declaration
    ;

node_declaration
    : node='place' id=ID (node_options)?
    | node='trans' id=ID (node_options)? (update_block)?
    | node=ID id=ID (node_options)? (update_block)?
    ;

arc_declaration
    : arctype=('arc'|'iarc'|'oarc'|'harc') srcName=ID 'to' destName=ID (node_options)?
    ;

reward_declaration
    : 'reward' id=ID expression
    ;

node_options
    : '(' option_list ')'
    ;

option_list
    : option_value (',' option_list)*
    ;

option_value
    : label_expression
    ;

label_expression
    : id=ID '=' expression
    ;

update_block
    : simple_block
    ;

simple_block
    : (NEWLINE)* '{' (simple)? (NEWLINE (simple)?)* '}'
    ;

simple
    : assign_expression
    | expression
    ;

// assign

assign_expression returns [int exprtype]
    : id=ID '=' expression {$exprtype = 1; }
    | ntoken_expression '=' expression {$exprtype = 2; }
    ;

// expression

expression returns [int nodetype]
    : op=('!'|'+'|'-') expression {$nodetype = 1; }
    | expression op=('*'|'/'|'div'|'mod') expression {$nodetype = 2; }
    | expression op=('+'|'-') expression {$nodetype = 3; }
    | expression op=('<'|'<='|'>'|'>=') expression {$nodetype = 4; }
    | expression op=('=='|'!=') expression {$nodetype = 5; }
    | expression op='&&' expression {$nodetype = 6; }
    | expression op='||' expression {$nodetype = 7; }
    | op='ifelse' '(' expression ',' expression ',' expression ')' {$nodetype = 8; }
    | function_expression {$nodetype = 9; }
    | ntoken_expression {$nodetype = 10; }
    | enable_expression {$nodetype = 14; }
    | literal_expression {$nodetype = 11; }
    | id=ID {$nodetype = 12; }
    |	'(' expression ')' {$nodetype = 13; }
    ;

// function_expression

function_expression
    : id=ID '(' function_args ')'
    ;

function_args
    : args_list
    | option_list
    ;
// arg

args_list
    : args_value (',' args_list)*
    ;

args_value
    : expression
    ;

// ntoken

ntoken_expression
    : '#' id=ID
    ;

// enable

enable_expression
    : '?' id=ID
    ;

// literal

literal_expression returns [int littype]
    : val=INT      {$littype = 1; }
    | val=FLOAT    {$littype = 2; }
    | val=LOGICAL  {$littype = 3; }
    | val=STRING   {$littype = 4; }
    ;

// TOKENS

LOGICAL: TRUE | FALSE ;

ID: CHAR (DIGIT+ | CHAR+ | '.')* ;

INT: DIGIT+ ;

FLOAT
    : DIGIT+ '.' (DIGIT+)? (EXPONENT)?
    | '.' (DIGIT+)? (EXPONENT)?
    | DIGIT+ EXPONENT
    ;

STRING : '"' ( ESCAPED_QUOTE | ~('\n'|'\r') )*? '"';

NEWLINE : [\r\n;EOF]+ ;

WS      : [ \t]+ -> skip ;

LINE_COMMENT
    : '//' ~[\r\n]* -> channel(HIDDEN)
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> channel(HIDDEN)
    ;

fragment
DIGIT: [0-9];

fragment
EXPONENT: [eE] ('+'|'-')? (DIGIT+)? ;

fragment
CHAR    : [a-zA-Z_] ;

fragment
TRUE    : [Tt][Rr][Uu][Ee] ;

fragment
FALSE   : [Ff][Aa][Ll][Ss][Ee] ;

fragment ESCAPED_QUOTE : '\\"';
