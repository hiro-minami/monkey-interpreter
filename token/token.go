package token

type TokenType string

// トークンを保持するための構造体
type Token struct {
	Type    TokenType
	Literal string
}

// トークンの種類を定義（定数）
const (
	// 未知の文字
	ILLEGAL = "ILLEGAL"
	// ファイルの最終行
	EOF = "EOF"

	// 識別子
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	// 括弧
	LEFTPAREN  = "("
	RIGHTPAREN = ")"
	LEFTBRACE  = "{"
	RIGHTBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
