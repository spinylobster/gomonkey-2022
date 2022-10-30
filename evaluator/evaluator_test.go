package evaluator_test

import (
	"gomonkey/evaluator"
	"gomonkey/lexer"
	"gomonkey/object"
	"gomonkey/parser"
	"testing"
)

// 私は「-」前置演算子のために新しいテスト関数を書くのではなく、このテストを拡張することにした。
// それには2つ理由がある。
// 第一に、前置の「-」演算子がサポートするオペランドは整数だけだからだ。
// 第二に、このテスト関数は全ての整数演算を含むように成長させ、期待する振る舞いを明確で整理された書き方で1つの場所にまとめておくためだ。
func TestIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// ast.IntegerLiteralなやつ
		{"5", 5},
		{"10", 10},

		// ast.PrefixExpressionなやつ
		{"-5", -5},
		{"-10", -10},

		// もうちょいテストケースを筋肉質にできると思うけど？
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		testIntegerObject(t, evaluated, tt.expected)
	}

}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	integerObj, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("obj is not *object.Integer. got=%T", obj)
		return false
	}

	if integerObj.Value != expected {
		t.Errorf("integerObj.Value not %d, got %d", expected, integerObj.Value)
		return false
	}

	return true
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return evaluator.Eval(program)
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	booleanObj, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("obj is not *object.Boolean. got=%[1]T, (%+[1]v)", obj)
		return false
	}

	if booleanObj.Value != expected {
		t.Errorf("booleanObj.Value is not %t, got %t", expected, booleanObj.Value)
		return false
	}

	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},

		// 特殊だぞ！
		{"!5", false},

		// 2連続のやつ
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},

		// Integerならとにかくfalseになる仕様。!0でも!1でもとにかくfalse
		{"!1", false},
		{"!0", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		testBooleanObject(t, evaluated, tt.expected)
	}
}
