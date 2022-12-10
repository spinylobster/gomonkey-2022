package evaluator

import (
	"fmt"
	"gomonkey/ast"
	"gomonkey/object"
	"gomonkey/token"
)

func quote(node ast.Node) object.Object {
	node = evalUnquoteCalls(node)
	return &object.Quote{Node: node}
}

func evalUnquoteCalls(quoted ast.Node) ast.Node {
	// to俺: ast.Modifyをまず呼びだしているからな！
	// 第2引数の 関数 はその後やで！
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		callExpr, ok := node.(*ast.CallExpression)
		if !ok {
			// CallExpressionじゃないなら何もしません
			// ex: quote(1 + 2)
			return node
		}

		// CallExpressionが確定！
		// もし、Functionが `unquote`なら、Modifyチャンス！
		if callExpr.Function.TokenLiteral() == "unquote" {
			// ex: unquote(1)だったら、ここだけ評価する
			// unquoteの引数は絶対1個なので、これでおk
			arg := callExpr.Arguments[0]

			// TODO: これは後でなんとかしましょう！
			嘘env := object.NewEnvironment()
			evaluated := Eval(arg, 嘘env)

			// object.Object -> ast.Node に変える
			// evaluatedの中身(INT前提)をつかまえて、IntegerLiteralに詰め直す
			intObj, _ := evaluated.(*object.Integer)
			intNode := &ast.IntegerLiteral{
				Token: token.Token{Type: "INT", Literal: fmt.Sprintf("%d", intObj.Value)},
				Value: intObj.Value,
			}
			return intNode
		}

		return node

	})
}
