package token

// 사용자 정의 자료형 TokenType
/*
	TokenType을 String으로 정의한 이유는 서로 다른 여러 값을 TokenType으로
	필요한 만큼 정의하여 사용할 수 있기 때문이다.
*/
type TokenType string

// 사용자 정의 구조체 Token
/*
	토큰이란 쉽게 분류할 수 있는 작은 자료구조이다.
	Token은 두개의 필드를 지닌다.
	1.Ty pe
		Type 필드는 토큰이 어떤 타입인지 나타내는 역활을 한다.
	2.Literal
	 	Literal 필디는 토큰의 값을 가지는 역활을 한다.

	예를 들면)
	`let x = 5;` 이러한 코드가 있다고 했을땐, 5가지 토큰이 생성된다.
	{LET, "let"},
	{IDENT, "x"},
	{ASSIGN, "="},
	{INT, "5"},
	{SEMICOLON, ";"},

*/
type Token struct {
	Type    TokenType
	Literal string
}

// Token Types
/*
	Token Type이란 이 토큰이 어떤 토큰인지 구별해주는 역활을 한다.
	마치 변수의 변수 타입과 유사한 개념이다. (이 변수가 어떤 자료형인가의 관점)

	요컨대 매우 많은 Token Types들이 존재하지만, 그 중 식별자의 말이 햇갈릴 수 있습니다.
	여기서 식별자란 예약어, 함수명, 변수명 등등을 포괄하는 의미이고
	밑에 IDENT tokenType은 사용자 정의 식별자를 나타냅니다.
*/
const (
	ILLEGAL = "ILLEGAL" // 어떤 토큰이나 문자를 렉서가 알 수 없다. 라는 의미로 사용됨

	EOF = "EOF" // 파서(Parser)에게 "이제 그만해도 좋다"라는 의미로 사용됨

	// 사용자 정의 식별자와 + Literal
	IDENT = "IDENT" // 사용자 정의 식별자를 나타냅니다.
	INT   = "INT"

	// 연산자
	ASSIGN = "="
	PLUS   = "+"

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Keywrods 맵
/*
	키워드에 대응하는 문자열을 tokenType이랑 매칭시켜논 자료구조 입니다.
*/
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookUpIdent 함수
// 매개변수 자료형:	 string
// 반환형 자료형 :  TokenType
/*
이 함수는 식별자 중 키워드인지 아닌지 판단하는 행위를 당담한다.

만약 string이 keywords에 대응되는 ToeknType이 있다면 해당 TokenType을 반환하고
없으면 반환하지 사용자 정의 식별자를 뜻하는 IDENT토큰을 반환한다.

*/
func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
