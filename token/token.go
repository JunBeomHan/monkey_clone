package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// TokenTypes
const (
	ILLEGAL = "ILLEGAL" // 어떤 토큰이나 문자를 렉서가 알 수 없다. 라는 의미로 사용됨

	EOF = "EOF" // 파서(Parser)에게 "이제 그만해도 좋다"라는 의미로 사용됨

	// Identifine + Literal
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
