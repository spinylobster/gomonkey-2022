package evaluator

import (
	"gomonkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			// argsは1個 && STRING => len(STRING) // 配列とかハッシュマップに対してのlenはまたあとでな！
			if len(args) != 1 {
				return newError("argument error: wrong number of arguments (given %d, expected %d)", len(args), 1)
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
}
