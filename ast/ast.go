package ast

import (
	"gomonkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string // ノードを文字列比較できると楽だから。Goの場合は、型が異なると直接比較できないからね。
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out strings.Builder

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // Identifierノード(token.IDENTではない！)
	Value Expression
}

func (ls *LetStatement) String() string {
	var out strings.Builder

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {
	panic("implement me")
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {
	panic("implement me")
}

type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var out strings.Builder

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) statementNode() {
	panic("implement me")
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// ExpressionStatement は 「1つの式」だけで構成される「文」
// x = 5; は let文だけど
// x + 5; は 式文
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (es *ExpressionStatement) expressionNode() {
	//TODO implement me
	panic("implement me")
}

func (es *ExpressionStatement) statementNode() {
	panic("implement me")
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.TokenLiteral()
}

type IntegerLiteral struct {
	Token token.Token // token.INT
	Value int64
}

func (il *IntegerLiteral) expressionNode() {
	panic("implement me")
}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token // 前置トークン、たとえば「！」
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {
	panic("implement me")
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out strings.Builder

	// たとえば、 (!5) になるってことですな
	// String() メソッドにおいて、演算子とそのオペランドとなる Right 内の式をわざと丸括弧で括っている。
	// このようにすることで、どのオペランドがどの演算子に属するのかがわかるようになる。
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // 演算子トークン
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {
	panic("implement me")
}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out strings.Builder

	// "(3 + 4)" みたいな表現にするってこと
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()

}

type Boolean struct {
	Token token.Token // token.TRUE | token.FALSE
	Value bool
}

func (b *Boolean) expressionNode() {
	panic("implement me")
}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

type IfExpression struct {
	Token       token.Token // token.IF
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) expressionNode() {
	panic("implement me")
}

func (ie *IfExpression) String() string {
	var out strings.Builder
	// if (cond) { consequence } else { alternative }
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
