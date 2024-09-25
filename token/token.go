package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	//ASSIGN   = "=" // Prophet don't have this
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "="
	NOT_EQ = "<>"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	THEN     = "THEN"
)

var keywords = map[string]TokenType{
	"fn":    FUNCTION, // Prophet don't have this
	"true":  TRUE,
	"false": FALSE,
	"IF":    IF,
	"If":    IF,
	"if":    IF,
	"iF":    IF,
	"ELSE":  ELSE,
	"Else":  ELSE,
	"THEN":  THEN,
	"Then":  THEN,
	"then":  THEN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
