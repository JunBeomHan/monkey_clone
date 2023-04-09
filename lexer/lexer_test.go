package lexer

import (
	"Monkey-clone/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for index, resTok := range tests {
		tok := l.NextToken()
		if tok.Type != resTok.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", index, resTok.expectedType, tok.Type)
		}
		if tok.Literal != resTok.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", index, resTok.expectedLiteral, tok.Literal)
		}
	}
}
