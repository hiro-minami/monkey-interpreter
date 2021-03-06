package lexer

import "monkey-interpreter/token"

type Lexer struct {
	input string
	// 入力における現在の位置（現在の文字）
	position int
	// これから読み込む位置（現在の次の文字）
	readPosition int
	// 現在検査中の文字
	ch byte
}

// 入力を受け取り字句解析器を返す
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// l.ch, l.position, l.readPositionを初期化
	l.readChar()
	return l
}

// 現在の文字を読み込み、inputの現在位置を進める
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ASCIIは0をNULLとして扱う
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	// 最後に読み取った位置を設定する
	l.position = l.readPosition
	l.readPosition += 1
}

// 次のトークンを取得する
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// 空白をスキップする処理を追加
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOTEQUAL, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LESSTHAN, l.ch)
	case '>':
		tok = newToken(token.GREATERTHAN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LEFTPAREN, l.ch)
	case ')':
		tok = newToken(token.RIGHTPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LEFTBRACE, l.ch)
	case '}':
		tok = newToken(token.RIGHTBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// readIdentifierの呼び出しの中でreadChar()を繰り返し読んでいる。
			// readPositionとpositionが現在の識別子の最後の文字を過ぎたところまで進んでいる
			// 上記のための、switch分の後でreadChar()を呼び出す必要がないため、このタイミングでreturnしている
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INTEGER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// 値を元にトークンを返す
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// このあたりInterface使ったらいい感じにできそうな気がする（宿題）
// 識別子を読んで非英字に到達するまで字句解析器を進める（
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 英字かどうかの判断
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// 空白をスキップする
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 次に来る値を覗き見する
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}