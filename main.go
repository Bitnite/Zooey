package main

import (
	"fmt"

	evaluator "github.com/ZooeyLang/Evaluator"
	"github.com/ZooeyLang/Lexer"
	object "github.com/ZooeyLang/Object"
	parser "github.com/ZooeyLang/Parser"
)

func main() {
	l := Lexer.NewLexer(`
		owo zap :=: 5; 

		for(owo i :=: 0; i <= 10; i++){
			zap :=: zap + 1
		}; 

		// ímplement prisma operators, berserker mode on
		prisma x :=: i(lo,ve) ~> zoo(e,y)
		
		zap`)

	parser := parser.New(l)
	program := parser.ParseProgram()

	environment := object.NewEnvironment()
	evaluator := evaluator.Eval(program, environment)

	fmt.Println(evaluator.Inspect())

	zap := 5
	for i := 0; i <= 10; i++ {
		zap += 1
	}
	fmt.Println(zap)
}
