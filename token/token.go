package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Token Types
const (
	ILLEGAL = "ILLEGAL" // 어떤 토큰이나 문자를 렉서가 알 수 없다. 라는 의미로 사용됨

	EOF = "EOF" // 파서(Parser)에게 "이제 그만해도 좋다"라는 의미로 사용됨

	// Identifine + Literal
	IDENT = "IDENT" // 사용자 정의 식별자를 나타냅니다.
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

// Keywrods
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

/*
func LookUpIdent(ident string) TokenType
- 함수의 쓰임새
이 함수는 식별자 중 키워드인지, 변수명인지, 함수명인지 판단하여 해당 TokenType을
반환하는 함수 입니다.

- 함수의 원리
먼저, 매개변수로 string 변수를 받고, keywords 맵에 해당 키에 대응하는 값이 있으면 TokenType을 반환합니다. (keywords안에는 monkey lang이 사용하는 키워드가 존재합니다.)
만약 없다면 사용자 정의 식별자를 나타내는 IDENT를 반환합니다.
*/
func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
