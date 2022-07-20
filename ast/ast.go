package ast

import "gomonkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program ノードは、Statementでもないし、Expressionでもないですね。
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return "" // ルートノードしかない場合ってこと(文が一切ない"プログラム"のとき)
	}
}

// LetStatement は 当然、Nodeだし、Statementですね。Expressionではないですね。
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // <identifier>
	Value Expression  // <expression>
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {
	panic("implement me")
}

// Identifier は Node です。さらに、monkeyの仕様では、 Identifier は Expression なのです！ もちろん Statement ではありません。
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {
	panic("implement me")
}
