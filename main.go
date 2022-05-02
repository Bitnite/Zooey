package main

import (
	"fmt"

	evaluator "github.com/ZooeyLang/Evaluator"
	"github.com/ZooeyLang/Lexer"
	object "github.com/ZooeyLang/Object"
	parser "github.com/ZooeyLang/Parser"
)

func main() {
	l := Lexer.NewLexer(`owo x :=: 0; for ( owo i :=: 0; i <= 100; i :=: i + 1) { x :=: x + 10; x}`)
	parser := parser.New(l)
	program := parser.ParseProgram()

	environment := object.NewEnvironment()
	evaluator := evaluator.Eval(program, environment)

	fmt.Println(evaluator.Inspect())
}
