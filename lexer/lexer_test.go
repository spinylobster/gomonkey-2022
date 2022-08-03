package lexer

import (
	"gomonkey/token"
	"testing"
)

func Test1文字トークンの字句解析(t *testing.T) {
	input := `(){}`

	tests := []struct {
		expectedToken   token.Type
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}
	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedToken {
			t.Fatalf("TokenTypeが違うよ。got=%q, want=%q", tok.Type, tt.expectedToken)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("TokenLiteralが違うよ。got=%s, want=%s", tok.Type, tt.expectedToken)
		}
	}
}
