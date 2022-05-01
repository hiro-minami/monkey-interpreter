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

// 定数に基づいたキーワード
var keywords = map[string]TokenType {
	"fn":  FUNCTION,
	"let": LET,
}

// keywordsテーブルをチェックして渡された引数がキーワードかどうかを判定
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}