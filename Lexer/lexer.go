package Lexer

import (
	"fmt"
	"strings"

	"github.com/ZooeyLang/Token"
)

type Lexer struct {
	input    string // Cadeia de caracters a ser recebida e tokenizada
	curChar  int    //cursor para posição atual na cadeia
	nextChar int    // posição seguinte ao cursor
	ch       byte   // char
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input} //Inicia um novo lexer com a cadeia de char passada
	lexer.readChar()              //Coloca o cursor na posição do primeiro caracter da cadeia

	return lexer
}

func (lexer *Lexer) readChar() {
	//Verifica se a proxima posição é o final da cadeia, atribuindo 0 em caso positivo

	if lexer.nextChar >= len(lexer.input) {
		lexer.ch = 0
	} else {
		//Caso contrário, temos de acessar o char atual

		lexer.ch = lexer.input[lexer.nextChar]
	}

	lexer.curChar = lexer.nextChar

	lexer.nextChar += 1
}

//Para cada elemento da cadeia de chars, analisamos o cursor e atribuimos um token a esse char
func (lexer *Lexer) NextToken() Token.Token {
	var tok Token.Token

	lexer.skipWhiteSpaces()

	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			tok = Token.Token{Type: Token.EQ, Literal: string(ch) + string(lexer.ch)}
		} else {
			tok = newToken(Token.ILLEGAL, lexer.ch)
		}

	case '+':
		tok = newToken(Token.PLUS, lexer.ch)

	case '"':
		tok.Type = Token.STRING
		tok.Literal = lexer.readString()

	case '-':
		tok = newToken(Token.MINUS, lexer.ch)

	case '!':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			tok = Token.Token{Type: Token.NOT_EQ, Literal: string(ch) + string(lexer.ch)}
		} else {
			tok = newToken(Token.BANG, lexer.ch)
		}
	case '&':
		if lexer.peekChar() == '&' {
			ch := lexer.ch
			lexer.readChar()
			tok = Token.Token{Type: Token.AND, Literal: string(ch) + string(lexer.ch)}
		}
	case '/':
		tok = newToken(Token.SLASH, lexer.ch)

	case '*':
		tok = newToken(Token.ASTERISK, lexer.ch)

	case ':':
		if lexer.peekChar() == '=' {
			aux := lexer.ch
			lexer.readChar()
			if lexer.peekChar() == ':' {
				ch := lexer.ch
				lexer.readChar()
				tok = Token.Token{Type: Token.ASSIGN, Literal: string(aux) + string(ch) + string(lexer.ch)}
			} else {
				tok = newToken(Token.ILLEGAL, lexer.ch)
			}
		} else {
			tok = newToken(Token.COLON, lexer.ch)
		}

	case '<':
		if lexer.peekChar() == '=' {

			ch := lexer.ch
			lexer.readChar()
			tok = Token.Token{Type: Token.LTE, Literal: string(ch) + string(lexer.ch)}
		} else {
			tok = newToken(Token.LT, lexer.ch)
		}

	case '>':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			tok = Token.Token{Type: Token.GTE, Literal: string(ch) + string(lexer.ch)}

		} else {
			tok = newToken(Token.GT, lexer.ch)
		}

	case ';':
		tok = newToken(Token.SEMICOLON, lexer.ch)

	case ',':
		tok = newToken(Token.DOT, lexer.ch)
	case '{':
		tok = newToken(Token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(Token.RBRACE, lexer.ch)
	case '(':
		tok = newToken(Token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(Token.RPAREN, lexer.ch)
	case '[':
		tok = newToken(Token.LBRACKET, lexer.ch)
	case ']':
		tok = newToken(Token.RBRACKET, lexer.ch)
	case 0:
		tok.Type = Token.EOF
		tok.Literal = ""

	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = Token.LookupIdent(tok.Literal)

			return tok
		} else if isDigit(lexer.ch) {
			tok.Literal = lexer.readNumber()
			if strings.Contains(tok.Literal, ".") == true {
				tok.Type = Token.FLOAT
			} else {
				tok.Type = Token.INT
			}

			return tok
		} else {
			tok = newToken(Token.ILLEGAL, lexer.ch)
		}

	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) readString() string {

	position := lexer.curChar + 1

	for {
		lexer.readChar()
		if lexer.ch == '"' || lexer.ch == 0 {
			break
		}
	}

	return lexer.input[position:lexer.curChar]

}

func (lexer *Lexer) skipWhiteSpaces() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func newToken(tokenType Token.TokenType, ch byte) Token.Token {
	return Token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) peekChar() byte {
	if lexer.nextChar >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.nextChar]
	}

}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) readIdentifier() string {
	firstPosition := lexer.curChar

	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[firstPosition:lexer.curChar]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

func (lexer *Lexer) readNumber() string {
	firstPosition := lexer.curChar
	var number string
	var midPosition int
	var dot bool

	for isDigit(lexer.ch) {
		if lexer.ch == '.' {
			number = lexer.input[firstPosition:lexer.curChar]
			number = fmt.Sprintf("%s", number)
			midPosition = lexer.curChar
			dot = true
		}
		lexer.readChar()
	}

	floatin := fmt.Sprintf("%s%s", number, lexer.input[midPosition:lexer.curChar])

	if dot {
		return floatin
	} else {
		return lexer.input[firstPosition:lexer.curChar]
	}

}
