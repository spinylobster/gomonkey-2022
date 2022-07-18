package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // 「未定義トークン」の定義も大事だね
	EOF     = "EOF"     //ファイル終端。構文解析器(パーサー)に終了をお知らせする

	IDENT = "IDENT" // 識別子: add, foobar, x, y, ...
	INT   = "INT"   // 整数: 3, 314

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
