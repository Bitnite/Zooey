package evaluator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	object "github.com/ZooeyLang/Object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(utf8.RuneCountInString(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"Benicio": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.String{Value: "Não foi possivel aniquilar lisbete, tente novamente mais tarde..."}
		},
	},
	"Stefano": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}

			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)

			array_assert := []string{"Verdadeira", "Falsa"}

			assert_1 := r1.Intn(1)

			s := fmt.Sprintf("A pessoa %s é com toda certeza -> %s\nE possui um total de %d de QI", args[0].Inspect(), array_assert[assert_1], r1.Intn(120))

			return &object.String{Value: s}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Argument to first must be array. got=%s, want=Array", args[0].Type())
			}

			arr := args[0].(*object.Array)

			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Argument to first must be array. got=%s, want=Array", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if len(arr.Elements) > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Argument to first must be array. got=%s, want=Array", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if len(arr.Elements) > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
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
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Argument to first must be array. got=%s, want=Array", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}

		},
	},
	"replace": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}
			if args[0].Type() != object.STRING {
				return newError("Argument to first must be array. got=%s, want=Array", args[0].Type())
			}

			to_be_changed := args[0].(*object.String)
			to_remove := args[1].(*object.String)
			to_put := args[2].(*object.String)

			return &object.String{Value: strings.Replace(to_be_changed.Value, to_remove.Value, to_put.Value, 3)}

		},
	},
	"plsShow": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			fmt.Println(args[0].Inspect())
			return nil
		},
	},
	"Strcomp": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			to_be_changed := args[0].(*object.String)
			to_remove := args[1].(*object.String)

			return &object.Boolean{Value: *to_be_changed == *to_remove}

		},
	},
	"String_Upcase": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			to_be_changed := args[0]
			return &object.String{Value: strings.ToUpper(to_be_changed.Inspect())}

		},
	},
	"String_Split": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			var objectArray []object.Object

			to_be_changed := args[0]

			x := strings.Split(to_be_changed.Inspect(), " ")

			for _, v := range x {
				objectArray = append(objectArray, &object.String{Value: v})
			}

			return &object.Array{Elements: objectArray}

		},
	},
}
