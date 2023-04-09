package lexer

import "monkey-clone/token"

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
	if l.readPosition >= len(l.input) {
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
	}
	l.readChar()
	return tok

}

// [Private]
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
