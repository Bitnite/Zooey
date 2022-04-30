package Lexer

import (
	"fmt"
	"strings"

	"github.com/ZooeyLang/Token"
)

type Lexer struct {
	input string // Cadeia de caracters a ser recebida e tokenizada
	cur   int    //cursor para posição atual na cadeia
	next  int    // posição seguinte ao cursor
	ch    byte   // char
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input} //Inicia um novo lexer com a cadeia de char passada
	lexer._readChar()             //Coloca o cursor na posição do primeiro caracter da cadeia

	return lexer
}

func (lexer *Lexer) _readChar() {
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

	lexer._skipWhiteSpaces()

	switch lexer.ch {

	case '=':
		if lexer._peekChar() == '=' {
			ch := lexer.ch
			lexer._readChar()
			tok = Token.Token{Type: Token.EQ, Literal: string(ch) + string(lexer.ch)}
		} else {
			tok = _newToken(Token.ILLEGAL, lexer.ch)
		}

	case '+':
		tok = _newToken(Token.PLUS, lexer.ch)

	case '"':
		tok.Type = Token.STRING
		tok.Literal = lexer._readString()

	case '-':
		tok = _newToken(Token.MINUS, lexer.ch)

	case '!':
		if lexer._peekChar() == '=' {
			ch := lexer.ch
			lexer._readChar()
			tok = Token.Token{Type: Token.NOT_EQ, Literal: string(ch) + string(lexer.ch)}
		} else {
			tok = _newToken(Token.BANG, lexer.ch)
		}

	case '/':
		tok = _newToken(Token.SLASH, lexer.ch)

	case '*':
		tok = _newToken(Token.ASTERISK, lexer.ch)

	case ':':
		if lexer._peekChar() == '=' {
			aux := lexer.ch
			lexer._readChar()
			if lexer._peekChar() == ':' {
				ch := lexer.ch
				lexer._readChar()
				tok = Token.Token{Type: Token.ASSIGN, Literal: string(aux) + string(ch) + string(lexer.ch)}
			} else {
				tok = _newToken(Token.ILLEGAL, lexer.ch) // isso ta certo??
			}
		} else {
			tok = _newToken(Token.COLON, lexer.ch)
		}

	case '<':
		if lexer._peekChar() == '=' {

			ch := lexer.ch
			lexer._readChar()
			tok = Token.Token{Type: Token.LTE, Literal: string(ch) + string(lexer.ch)}
		} else {
			tok = _newToken(Token.LT, lexer.ch)
		}

	case '>':
		if lexer._peekChar() == '=' {
			ch := lexer.ch
			lexer._readChar()
			tok = Token.Token{Type: Token.GTE, Literal: string(ch) + string(lexer.ch)}

		} else {
			tok = _newToken(Token.GT, lexer.ch)
		}

	case ';':
		tok = _newToken(Token.SEMICOLON, lexer.ch)

	case ',':
		tok = _newToken(Token.COMMA, lexer.ch)

	case '{':
		tok = _newToken(Token.LBRACE, lexer.ch)
	case '}':
		tok = _newToken(Token.RBRACE, lexer.ch)
	case '(':
		tok = _newToken(Token.LPAREN, lexer.ch)
	case ')':
		tok = _newToken(Token.RPAREN, lexer.ch)

	case '[':
		tok = _newToken(Token.LBRACKET, lexer.ch)
	case ']':
		tok = _newToken(Token.RBRACKET, lexer.ch)

	case 0:
		tok.Type = Token.EOF
		tok.Literal = ""

	default:
		if _isLetter(lexer.ch) {
			tok.Type = Token.LookupIdent(tok.Literal)
			tok.Literal = lexer._readIdentifier()

			return tok
		} else if _isDigit(lexer.ch) {
			tok.Literal = lexer._readNumber()
			if strings.Contains(tok.Literal, ".") == true {
				tok.Type = Token.FLOAT
			} else {
				tok.Type = Token.INT
			}

			return tok
		} else {
			tok = _newToken(Token.ILLEGAL, lexer.ch)
		}

	}

	lexer._readChar()
	return tok
}

func (lexer *Lexer) _readString() string {

	position := lexer.cur + 1

	for {
		lexer._readChar()
		if lexer.ch == '"' || lexer.ch == 0 {
			break
		}
	}

	return lexer.input[position:lexer.cur]

}

func (lexer *Lexer) _skipWhiteSpaces() {

	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer._readChar()
	}

}

func _newToken(tokenType Token.TokenType, ch byte) Token.Token {
	return Token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) _peekChar() byte {

	if lexer.next > len(lexer.input) {
		return 0
	} else {

		return lexer.input[lexer.next]

	}

}

func _isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) _readIdentifier() string {
	first_position := lexer.cur

	for _isLetter(lexer.ch) {
		lexer._readChar()
	}

	return lexer.input[first_position:lexer.cur]
}

func _isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

func (lexer *Lexer) _readNumber() string {
	first_position := lexer.cur
	var number string
	var position_middle int
	var i_have_dot bool

	for _isDigit(lexer.ch) {
		if lexer.ch == '.' {
			number = lexer.input[first_position:lexer.cur]
			number = fmt.Sprintf("%s", number)
			position_middle = lexer.cur
			i_have_dot = true
		}
		lexer._readChar()
	}

	floatin := fmt.Sprintf("%s%s", number, lexer.input[position_middle:lexer.cur])

	if i_have_dot {
		return floatin
	} else {
		return lexer.input[first_position:lexer.cur]
	}

}
