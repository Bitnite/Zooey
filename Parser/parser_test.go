package parser

import (
	"log"
	"testing"

	"github.com/ZooeyLang/Lexer"
	token "github.com/ZooeyLang/Token"
)

func TestParser_x(t *testing.T) {
	type test struct {
		name    string
		input   string
		want    []token.Token
		wantErr bool
	}
	l := Lexer.New(`x := 5; for ( owo i :=: 10; i >= 0; i :=: i - 1) { x := x - 10; }`)
	parser := New(l)
	program := parser.ParseProgram()

	for _, e := range program.Statements {
		log.Println(e.String())
	}

}
