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

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	BANG = "!"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	// 文字列っぽいものは、キーワードかもしれないし、
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT //識別子かもしれないね
}
