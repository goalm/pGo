package lexer

import (
	"fmt"
	"github.com/goalm/pGo/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
IF t>0 THEN
	IF SURP_RES_IND <> 1 THEN
		AE_AS_IF_S25(t) + SH_EXCAP_IF_SURP(t)
	ELSE
    	C_MATHRES_IF(t)
ELSE
	0`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.IF, "IF"},
		{token.IDENT, "t"},
		{token.GT, ">"},
		{token.INT, "0"},
		{token.THEN, "THEN"},
		{token.IF, "IF"},
		{token.IDENT, "SURP_RES_IND"},
		{token.NOT_EQ, "<>"},
		{token.INT, "1"},
		{token.THEN, "THEN"},
		{token.IDENT, "AE_AS_IF_S25"},
		{token.LPAREN, "("},
		{token.IDENT, "t"},
		{token.RPAREN, ")"},
		{token.PLUS, "+"},
		{token.IDENT, "SH_EXCAP_IF_SURP"},
		{token.LPAREN, "("},
		{token.IDENT, "t"},
		{token.RPAREN, ")"},
		{token.ELSE, "ELSE"},
		{token.IDENT, "C_MATHRES_IF"},
		{token.LPAREN, "("},
		{token.IDENT, "t"},
		{token.RPAREN, ")"},
		{token.ELSE, "ELSE"},
		{token.INT, "0"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println(tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
