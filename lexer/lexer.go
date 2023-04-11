package lexer

import (
	"monkey_clone/token"
)

/*
	함수:
		func name
		public? private?
		리시버
		매개변수
		반환값
*/

// [Public] Lexer struct 생성자
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	/*
		res filed value:
			readPosition = 1
			Postion = 0
			ch = input[0]
	*/
	return l
}

// [Public]
type Lexer struct {
	input        string
	position     int  // 입력에서 현재 위치(현재 문자를 가리킨다.)
	readPosition int  // 입력에서 현재 읽는 위치 (현재 문자의 다음을 가리킨다.)
	ch           byte // 현재 조사하고 있는 문자
	/*
		byte를 사용하는 한다. -> 아스키만 지원한다.
		만약 유니코드를 사용할 시 -> byet(x) rune(o)
	*/
}

// [Private] 다음 문자의 접근하고 저장하는 행위
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) /*EOF라면?*/ {
		l.ch = 0 // l.ch = NUL -> (아직 아무것도 읽지 않는 상태 or 파일의 끝(EOF))
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	/*
		l.string[positoin] == l.ch
	*/
}

// [Public]
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case 0:
		// newToken 함수를 사용하지 않은 이유는 tok.Literal 값에 ""를 대입해주기 위해서.
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// 식별자를 확인. (여기서 식별자란 키워드, 함수명, 변수명이 포함된다..? -> TODO:: Seach해보기)
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifer()
			// Token의 타입은 LookUpIdent함수로 실행하여 얻어온다. 그 이유는 식별자는 종류가 여러개이기 때문이다.
			// LookUpIdent 함수는 token/token.go에 있음
			tok.Type = token.LookUpIdent(tok.Literal)
			// 여기서 중요!! 그냥 이렇게 작성한 경우 식별자가 연달아 나온 경우
			// 공백을 처리하지 못하고 ILLEGAL 토큰을 반환한다. 그러므로 공백을 지나가는 함수를 정의해줘야 한다.
			return tok
		} else if isDigit(l.ch) { // 숫자 처리
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// 알수 없는 오류 처리
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
	// 다른 파서에서는 eatWhitespace, comsumeWhitespace 혹은 아예 다른 이름으로 불린다.
}

// [Private]
// 문자 그대로 동작하는 함수이다.
// 식별자 하나를 읽어 들이고 렉서의 position를 저장한 뒤 글자가 아닐 떄까지 돌린 postition까지 식별자로 판단한다.
func (l *Lexer) readIdentifer() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // position ~ (l.position - 1)
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	/*
		실제로는 2진수, 8진수, 16진수를 생각해야 한다.
	*/
	return ch >= '0' && ch <= '9'
}

// [Private]
// 파싱의 전반적인 과정에서 중요한 역확을 한다. 왜냐하면
// ch == '_' 여기서 _는 식별자를 허용한다는 말이기 때문이다.
// 만약 !를 식별자로 사용하고 싶으면 아래의 코드에 ch == '!'만 삽입하면 된다.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// [Private]
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
