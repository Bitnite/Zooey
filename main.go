package main

import (
	"fmt"

	evaluator "github.com/ZooeyLang/Evaluator"
	lexer "github.com/ZooeyLang/Lexer"
	object "github.com/ZooeyLang/Object"
	parser "github.com/ZooeyLang/Parser"
)

func main() {

	env := object.NewEnvironment()
	l := lexer.New(`
	fn passouDeAno(nota){
		if nota > 7 {
			return true
		}
		return false
	}

	owo notaUm :=: 7; 
	owo notaDois :=: 8; 
	owo notaTres :=: 5; 
	owo notaQuatro :=: 4; 

	owo media :=: (notaUm + notaDois + notaTres + notaQuatro) / 4;
	show(media)
	show(passouDeAno(media))

	for(owo i :=: 0; i <= 10; i++){
		media :=: media + 10
	}

	show(media)

	owo testeEscopo :=: "eu valho porcaria ninhuma..."

	fn valorizador(testeEscopo, enzo){
	 testeEscopo :=: "finalmente reconheceram meu valor..."
	 show(enzo);
	}
	valorizador("oi", "aula")
	show(testeEscopo)


	owo verdadeiro :=: true 
	owo falso :=: false

	`)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println(p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated)
	}
}
