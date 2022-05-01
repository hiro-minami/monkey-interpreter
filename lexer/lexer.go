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

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case'(':
		tok = newToken(token.LEFTPAREN, l.ch)
	case ')':
		tok = newToken(token.RIGHTPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LEFTBRACE, l.ch)
	case '}':
		tok = newToken(token.RIGHTBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

// 値を元にトークンを返す
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}