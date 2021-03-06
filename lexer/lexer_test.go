package lexer

import (
	"monkey-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	10 == 10;
	10 != 9;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LEFTPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RIGHTPAREN, ")"},
		{token.LEFTBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LEFTPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RIGHTPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "5"},
		{token.LESSTHAN, "<"},
		{token.INTEGER, "10"},
		{token.GREATERTHAN, ">"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LEFTPAREN, "("},
		{token.INTEGER, "5"},
		{token.LESSTHAN, "<"},
		{token.INTEGER, "10"},
		{token.RIGHTPAREN, ")"},
		{token.LEFTBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRACE, "}"},
		{token.ELSE, "else"},
		{token.LEFTBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRACE, "}"},
		{token.INTEGER, "10"},
		{token.EQUAL, "=="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "10"},
		{token.NOTEQUAL, "!="},
		{token.INTEGER, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		// ??????????????????????????????????????????????????????????????????
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		// ????????????????????????????????????????????????
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
