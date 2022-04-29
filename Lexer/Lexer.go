package Lexer

import "github.com/ZooeyLang/Token"

type Lexer struct {
	input string // Cadeia de caracters a ser recebida e tokenizada
	cur   int    //cursor para posição atual na cadeia
	next  int    // posição seguinte ao cursor
	ch    byte   // char
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input} //Inicia um novo lexer com a cadeia de char passada
	lexer.readChar()              //Coloca o cursor na posição do primeiro caracter da cadeia

	return lexer
}

func (lexer *Lexer) readChar() {
	//Verifica se a proxima posição é o final da cadeia, atribuindo 0 em caso positivo

	if lexer.next >= len(lexer.input) {
		lexer.ch = 0
	} else {
		//Caso contrário, temos de acessar o char atual

		lexer.ch = lexer.input[lexer.next]
	}

	lexer.cur = lexer.next

	lexer.next += 1
}

//Para cada elemento da cadeia de chars, analisamos o cursor e atribuimos um token a esse char

func (lexer *Lexer) NextToken() Token.Token {

	var tok Token.Token

}

func (lexer *Lexer) _skipWhiteSpaces() {

	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}

}
