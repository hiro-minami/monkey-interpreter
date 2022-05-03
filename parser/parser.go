package parser

import (
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// 2つトークンを読み込む
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
