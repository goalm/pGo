package token

type Type string
type Token struct {
	Type    Type
	Literal string
}

const (
	DEFALUT = "DEFAULT"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	COMMENT = "COMMENT"

	// Identifiers + literals
	CONST   = "CONST"
	IDENT   = "IDENT"   // add, foobar, x, y, ...
	INTEGER = "INTEGER" // 1343456
	FLOAT   = "FLOAT"   // 1343456.123
	STRING  = "STRING"

	// Operators
	ASSIGN   = ":="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	MULTIPLY = "*"
	DIVIDE   = "/"

	LT     = "<"
	LTE    = "<="
	GT     = ">"
	GTE    = ">="
	EQ     = "="
	NOT_EQ = "<>"

	// Logical Operators
	NOT = "NOT"
	AND = "AND"
	OR  = "OR"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	DOT       = "."

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	//T        = "T"

	// Keywords
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	THEN     = "THEN"
	FOR      = "FOR"

	// Prophet functions
	// math functions
	ABS         = "ABS"
	DIV         = "DIV"
	EXP         = "EXP"
	FRACT       = "FRACT"
	GROUP_AFTER = "GROUP_AFTER"
	INT         = "INT"
	INTERP      = "INTERP"
	LN          = "LN"
	LOG         = "LOG"
	MAX         = "MAX"
	MAXT        = "MAXT"
	MIN         = "MIN"
	MINT        = "MINT"
	MOD         = "MOD"
	NO_CALC     = "NO_CALC"
	PROD        = "PROD"
	ROUND       = "ROUND"
	ROUND_DOWN  = "ROUND_DOWN"
	ROUND_NEAR  = "ROUND_NEAR"
	ROUND_UP    = "ROUND_UP"
	SUM         = "SUM"

	// array functions
	ARRAY_INITIAZE = "ARRAY_INITIAZE"
	ARRAY_LBOUND   = "ARRAY_LBOUND"
	ARRAY_MAX      = "ARRAY_MAX"
	ARRAY_MIN      = "ARRAY_MIN"
	ARRARY_PROD    = "ARRARY_PROD"
	ARRAY_SIZE     = "ARRAY_SIZE"
	ARRAY_SUM      = "ARRAY_SUM"
	ARRAY_UBOUND   = "ARRAY_UBOUND"

	// conversion functions
	ENUM_SIZE    = "ENUM_SIZE"
	ENUM_TO_INT  = "ENUM_TO_INT"
	ENUM_TO_TEXT = "ENUM_TO_TEXT"
	INT_TO_ENUM  = "INT_TO_ENUM"
	NUM_TO_TEXT  = "NUM_TO_TEXT"
	TEXT_TO_ENUM = "TEXT_TO_ENUM"
	TEXT_TO_NUM  = "TEXT_TO_NUM"

	// dynamic functions
	BUY_ASSET           = "BUY_ASSET"
	CURRENT_START_MONTH = "CURRENT_START_MONTH"
	CURRENT_START_YEAR  = "CURRENT_START_YEAR"
	FIRST_CALC          = "FIRST_CALC"
	PROP_BUY_ASSET      = "PROP_BUY_ASSET"
	SELL_ASSET          = "SELL_ASSET"

	// financial functions
	BEM       = "BEM"
	BLACK_SCH = "BLACK_SCH"
	IRR       = "IRR"
	MONINT    = "MONINT"

	// logical functions
	MULT = "MULT"

	// message functions
	ERROR         = "ERROR"
	WARNING       = "WARNING"
	PRINT         = "PRINT"
	PRINT_TO_FILE = "PRINT_TO_FILE"
)

var keywords = map[string]Type{
	"fn": FUNCTION, // Prophet don't have this
	//"t":                   T,
	//"T":                   T,
	"true":                TRUE,
	"false":               FALSE,
	"IF":                  IF,
	"If":                  IF,
	"if":                  IF,
	"iF":                  IF,
	"ELSE":                ELSE,
	"Else":                ELSE,
	"else":                ELSE,
	"THEN":                THEN,
	"Then":                THEN,
	"then":                THEN,
	"AND":                 AND,
	"And":                 AND,
	"and":                 AND,
	"OR":                  OR,
	"Or":                  OR,
	"or":                  OR,
	"NOT":                 NOT,
	"Not":                 NOT,
	"not":                 NOT,
	"FOR":                 FOR,
	"For":                 FOR,
	"for":                 FOR,
	"ABS":                 ABS,
	"DIV":                 DIV,
	"EXP":                 EXP,
	"FRACT":               FRACT,
	"GROUP_AFTER":         GROUP_AFTER,
	"INT":                 INT,
	"INTERP":              INTERP,
	"LN":                  LN,
	"LOG":                 LOG,
	"MAX":                 MAX,
	"MAXT":                MAXT,
	"MIN":                 MIN,
	"MINT":                MINT,
	"MOD":                 MOD,
	"NO_CALC":             NO_CALC,
	"PROD":                PROD,
	"ROUND":               ROUND,
	"ROUND_DOWN":          ROUND_DOWN,
	"ROUND_NEAR":          ROUND_NEAR,
	"ROUND_UP":            ROUND_UP,
	"SUM":                 SUM,
	"ARRAY_INITIAZE":      ARRAY_INITIAZE,
	"ARRAY_LBOUND":        ARRAY_LBOUND,
	"ARRAY_MAX":           ARRAY_MAX,
	"ARRAY_MIN":           ARRAY_MIN,
	"ARRARY_PROD":         ARRARY_PROD,
	"ARRAY_SIZE":          ARRAY_SIZE,
	"ARRAY_SUM":           ARRAY_SUM,
	"ARRAY_UBOUND":        ARRAY_UBOUND,
	"ENUM_SIZE":           ENUM_SIZE,
	"ENUM_TO_INT":         ENUM_TO_INT,
	"ENUM_TO_TEXT":        ENUM_TO_TEXT,
	"INT_TO_ENUM":         INT_TO_ENUM,
	"NUM_TO_TEXT":         NUM_TO_TEXT,
	"TEXT_TO_ENUM":        TEXT_TO_ENUM,
	"TEXT_TO_NUM":         TEXT_TO_NUM,
	"BUY_ASSET":           BUY_ASSET,
	"CURRENT_START_MONTH": CURRENT_START_MONTH,
	"CURRENT_START_YEAR":  CURRENT_START_YEAR,
	"FIRST_CALC":          FIRST_CALC,
	"PROP_BUY_ASSET":      PROP_BUY_ASSET,
	"SELL_ASSET":          SELL_ASSET,
	"BEM":                 BEM,
	"BLACK_SCH":           BLACK_SCH,
	"IRR":                 IRR,
	"MONINT":              MONINT,
	"MULT":                MULT,
	"ERROR":               ERROR,
	"WARNING":             WARNING,
	"PRINT":               PRINT,
	"PRINT_TO_FILE":       PRINT_TO_FILE,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
