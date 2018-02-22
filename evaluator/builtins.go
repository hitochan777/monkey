package evaluator

import (
	"github.com/hitochan777/monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			arg, ok := args[0].(*object.Array)

			if !ok {
				return newError("argument to `first` must be ARRAY. got %s", args[0].Type())
			}

			if len(arg.Elements) > 0 {
				return arg.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			arg, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `last` must be ARRAY. got %s", args[0].Type())
			}

			length := len(arg.Elements)
			if length > 0 {
				return arg.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			arg, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `rest` must be ARRAY. got %s", args[0].Type())
			}

			length := len(arg.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arg.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("first argument to `push` must be ARRAY. got %s", args[0].Type())
			}

			length := len(array.Elements)
			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, array.Elements)
			newElements[length] = args[1]
			return &object.Array{Elements: newElements}
		},
	},
}
