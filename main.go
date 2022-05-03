package main

import (
	"fmt"

	evaluator "github.com/ZooeyLang/Evaluator"
	"github.com/ZooeyLang/Lexer"
	object "github.com/ZooeyLang/Object"
	parser "github.com/ZooeyLang/Parser"
)

func main() {
	// ímplement prisma operators
	//prisma x :=: i(lo,ve) ~> zoo(e,y)

	l := Lexer.NewLexer(`
		owo zap :=: 5; 

		for(owo i :=: 0; i <= 10; i++){
			zap :=: zap + 1
		}; 

		zap`)

	parser := parser.New(l)
	program := parser.ParseProgram()

	environment := object.NewEnvironment()
	evaluator := evaluator.Eval(program, environment)

	fmt.Println(evaluator.Inspect())
}
