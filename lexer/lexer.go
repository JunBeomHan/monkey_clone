package lexer

import (
	"monkey_clone/token"
)

// 사용자 정의 구조체 Lexer
/*
	Lexer란 소스코드를 토큰 열로 변화시켜 주는 작업시켜주는 친구이다. 이 작업을 어휘 분석(Lexical analysis) 혹은 줄여서 렉싱(Lexing) 이라고 한다.
	Lexer의 필드에는 input, position, readPosition, ch가 존재한다.
	input: Lexer_test.go 에서 볼수 있듯이 소스코드를 넘겨주는 것을 확인할 수 있다. 즉 input 소스코드를 뜻한다.
	position: 소스코드를 읽을때 한자 한자 읽게 되는데, position은 현재 문자의 위치를 뜻한다.
	readPosition: readChar() 함수를 보면, readPosition이 항상 앞을 나서며 position을 이끈다.
				  입력 문자열을 가리기킄 포인터가 두개 인 이유는 다음 처리 대상을 알아내려면 다음 문자를 미리 살펴봄 동시에 현재 문자를 보존하기 위해서 이다.
				  readPosition은 항상 다음 문자를 가리키고, position은 현재 문자를 가리킨다.
	ch: input[position] 위치에 있는 문자 값이다. 이 문자를 보고 식별자, 연산자 등등 판단한다.
		여기서 중요한 점은 자료형이 byte인 점이다. byte로 정의했다는 뜻은 아스키코드만 사용하다는 점이다. 그렇기 때문에 매우 중요하다 볼 수 있다.
		만약 유니코드(한글 코딩)을 가능케 하고 싶다면, rune자료형을 사용하면 된다.
	*/
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

// New 함수 
// 매개변수 자료형 string
// 반환형 자료형  *Lexer
/*
이 함수는 Lexer를 생성해주는 행위를 당담한다.

먼저 string을 매개변수로 받은 뒤 Lexer의 input값에 대입 해주고, 렉서를 생성한 뒤 포인터 값을 l에 저장시킨다.
그런 뒤 l.readChar()를 실행시켜 코드를 읽을 준비를 한다.
그런 뒤 l를 반환한다.
*/
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	/*
		현재 l의 필드 주소값 상태
			readPosition = 1
			Postion = 0
			ch = input[0]
	*/
	return l
}

// readChar 함수 (l의 포인터 리시버)
// 매개변수 자료형: 없음
// 반환형 자료형: 없음
/*
	이 함수는 소스코드를 문자 하나하나를 읽어 ch에 저장하는 행위를 당담한다.

	l.readPostiton을 이용해 소스코드의 끝을 의미하는 EOF를 판단한다. 끝이면 l.ch에 끝을 뜻하는(EOF) 0을 대입해준다.
	끝이 아니라면 l.input[readPostiton]의 값을 담아준다.
	l.postion = l.readPostion이 되고
	l.readPostion += 1이 된다.

	이렇게 함으로써 한칸 한칸 해석을 할 수 있다.
*/
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

// NextToken 함수 (l의 포인터 리시버)
// 매개변수 자료형: 없음
// 반환형 자료형: TToken(사용자 정의 구조체 Token)
/*
	이 함수는 l.ch의 값을 보고 그에 알맞는 토큰을 반환하는 행위를 당담한다.

	먼저 l.skopWhitespace() 함수를 실행시켜 공백을 무시해준다. 
	이 작업이 중요한게, 만약 공백이 중심이 언어를 설계하면 공백을 무시하면 안돼고 
	토큰을 추출해야 한다.

	그 다음 연산자에 관한 토큰들을 추출한다. 

	만약 파일의 끝을 뜻하는 EOF가 나오면 Literal 값에 ""를 대입해주고 Type에는 EOF를 대입하고 반환해준다.


	근데 식별자는 어떻게 처리할까?
	이 코드는 default: 부분에 있다.

	먼저 isLetter함수를 이용헤 l.ch가 글자 인지 확인한다. (여기서 글자와 문자는 같은 뜻으로 사용된다.)
	여기서 중요한 점은 l.ch가 현재 _도 글자에 포함하고 있다는 점이다. 만약 !와 같은 이런 문자를 사용한다면 !처리문을 작성해주어야 한다.

	암튼 글자이면 식별자로 판단하기 떄문에 readIdentifer() 함수를 이용해 식별자를 따온다.
	그런 뒤 type에는 이 값을 매개변수로 전달해 해당 식별자에 대응되는 Type을 가진다.
	그리고 반환한다.

	어! 그런데 만약 문자가 아니고 숫자이면? 
	마찬가지로 isDigit함수를 실행해 문자 처리를 해주면 된다.
	하지만 여기서 중요한점은 정수, 실수, 16진수, 2진수 등 숫자를 여러가지 형태를 표현하는 케이스도 포함해야 한다.
	우리는 10진수 정수만 다룬다.

	어! 숫자도 아니고 식별자도 아니면 (공백과 같은것 혹은 !과 같은) 어떤 토큰이나 문자를 렉서가 알수 없다는 뜻을 가지고 있는 ILLEGAL토큰을 대입한 후 반환한다.

	아!아!아! 중요한 점이 있다. 위 세 가지 케이스를 제외한 그냥 연산자는 readChar()를 통해서 한글자 한글자 해석해야간다!

	참고로 토큰 생성은 newToken함수를 이용한다.
*/
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

// skipWhitespace 함수 (l의 포인터 리시버)
// 매개변수 자료형: 없음
// 반환 자료형: 없음
/*
	skipWhitespace함수는 l.ch가 공백인지 아닌지 판단하여 공백을 뛰어넘는 행위를 당담한다.
	다른 파서에서는 eatWhitespace, comsumeWhitespace 혹은 아예 다른 이름으로 불린다.
*/
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
	
}

// newToken 함수 
// 매개변수 자료형: TokenType, byte
// 반환 자료형: Token
/*
	newToken 토큰을 생성하고 반환하는 행위를 당담한다.
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}


// readIdentifer 함수 (l의 포인터 리시버)
// 매개변수 자료형: 없음
// 반환 자료형: string
/*
	readIdentifer함수는 식별자를 추출해 반환해주는 행위를 담당하고 있다.

	현재 position을 저장하고, isLetter함수를 실행해 문자가 아닌 값을 만났을 때까지 한칸 한칸 간다.
	그런 뒤 슬라이싱 기능을 사용하면 식별자를 추출할 수 있다.
*/

func (l *Lexer) readIdentifer() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // position ~ (l.position - 1)
}

// isLetter 함수 
// 매개변수 자료형: byte
// 반환 자료형: bool
/*
	isLetter 함수는 byte를 매개변수로 받아 이 문자가 글자인지 판단하여 참과 거짓을 반환하는 행위를 당담한다.
	여기서 중요한 점은 과연 어디까지 문자로 포함할 것인가이다. Monkey-lang은 _을 포함하고 있지만 다른 언어에서는 포함을 안할 수도 있다.
	그렇기에 잘 판다해야 한다.
*/
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}


// readIdentifer 함수 (l의 포인터 리시버)
// 매개변수 자료형: 없음
// 반환형 자료형: string	
/*
	readIdentifer 함수와 동일하게 숫자를 추출하여 반환해주는 행위를 당담한다.
	원리는 readIdentifer와 같다.
*/
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isDigit 함수
// 매개변수 자료형: byte
// 반환형 자료형: bool
/*
	isDigit함수는 byte를 매개변수로 받아 숫자인지 판단하여 참과 거짓을 판단하는 행위를 당담한다.
	원리는 isLetter() 함수와 같다.
*/
func isDigit(ch byte) bool {
	/*
		실제로는 2진수, 8진수, 16진수를 생각해야 한다.
	*/
	return ch >= '0' && ch <= '9'
}


