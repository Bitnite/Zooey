package Lexer

import (
	"testing"

	"github.com/ZooeyLang/Token"
	"github.com/stretchr/testify/assert"
)

func TestLexer_NextToken(t *testing.T) {
	type test struct {
		input   string
		want    []Token.Token
		wantErr bool
	}

	tests := []test{
		{
			input: "owo teste :=: 5;",
			want: []Token.Token{
				{Type: Token.OwO, Literal: "owo"},
				{Type: Token.IDENT, Literal: "teste"},
				{Type: Token.ASSIGN, Literal: ":=:"},
				{Type: Token.INT, Literal: "5"},
				{Type: Token.SEMICOLON, Literal: ";"},
			},
			wantErr: false,
		},
		{
			input: "if 5 > 2 { return true }",
			want: []Token.Token{
				{Type: Token.IF, Literal: "if"},
				{Type: Token.INT, Literal: "5"},
				{Type: Token.GT, Literal: ">"},
				{Type: Token.INT, Literal: "2"},
				{Type: Token.LBRACE, Literal: "{"},
				{Type: Token.RETURN, Literal: "return"},
				{Type: Token.TRUE, Literal: "true"},
				{Type: Token.RBRACE, Literal: "}"},
			},
			wantErr: false,
		},
		{
			input: "if 2 == 2 { return true }",
			want: []Token.Token{
				{Type: Token.IF, Literal: "if"},
				{Type: Token.INT, Literal: "2"},
				{Type: Token.EQ, Literal: "=="},
				{Type: Token.INT, Literal: "2"},
				{Type: Token.LBRACE, Literal: "{"},
				{Type: Token.RETURN, Literal: "return"},
				{Type: Token.TRUE, Literal: "true"},
				{Type: Token.RBRACE, Literal: "}"},
			},
			wantErr: false,
		},
		{
			input: "if 5 != 2 { return true }",
			want: []Token.Token{
				{Type: Token.IF, Literal: "if"},
				{Type: Token.INT, Literal: "5"},
				{Type: Token.NOT_EQ, Literal: "!="},
				{Type: Token.INT, Literal: "2"},
				{Type: Token.LBRACE, Literal: "{"},
				{Type: Token.RETURN, Literal: "return"},
				{Type: Token.TRUE, Literal: "true"},
				{Type: Token.RBRACE, Literal: "}"},
			},
			wantErr: false,
		},
		{
			input: "owo myFloat :=: 5.25",
			want: []Token.Token{
				{Type: Token.OwO, Literal: "owo"},
				{Type: Token.IDENT, Literal: "myFloat"},
				{Type: Token.ASSIGN, Literal: ":=:"},
				{Type: Token.FLOAT, Literal: "5.25"},
			},
			wantErr: false,
		},
		{
			input: "owo myInt :=: 5",
			want: []Token.Token{
				{Type: Token.OwO, Literal: "owo"},
				{Type: Token.IDENT, Literal: "myInt"},
				{Type: Token.ASSIGN, Literal: ":=:"},
				{Type: Token.INT, Literal: "5"},
			},
			wantErr: false,
		},
		{
			input: `owo x 5 "xx" 10.25 true false :=: + - ! * / > < >= <= == != ; : fn && for while`,
			want: []Token.Token{
				{Type: Token.OwO, Literal: "owo"},
				{Type: Token.IDENT, Literal: "x"},
				{Type: Token.INT, Literal: "5"},
				{Type: Token.STRING, Literal: "xx"},
				{Type: Token.FLOAT, Literal: "10.25"},
				{Type: Token.TRUE, Literal: "true"},
				{Type: Token.FALSE, Literal: "false"},
				{Type: Token.ASSIGN, Literal: ":=:"},
				{Type: Token.PLUS, Literal: "+"},
				{Type: Token.MINUS, Literal: "-"},
				{Type: Token.BANG, Literal: "!"},
				{Type: Token.ASTERISK, Literal: "*"},
				{Type: Token.SLASH, Literal: "/"},
				{Type: Token.GT, Literal: ">"},
				{Type: Token.LT, Literal: "<"},
				{Type: Token.GTE, Literal: ">="},
				{Type: Token.LTE, Literal: "<="},
				{Type: Token.EQ, Literal: "=="},
				{Type: Token.NOT_EQ, Literal: "!="},
				{Type: Token.SEMICOLON, Literal: ";"},
				{Type: Token.COLON, Literal: ":"},
				{Type: Token.FN, Literal: "fn"},
				{Type: Token.AND, Literal: "&&"},
				{Type: Token.FOR, Literal: "for"},
				{Type: Token.WHILE, Literal: "while"},
			},
			wantErr: false,
		},
		{
			input: ":=",
			want: []Token.Token{
				{Type: Token.ILLEGAL, Literal: "="},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		l := NewLexer(tc.input)

		var tokenEOF Token.Token
		tokenEOF.Type = Token.EOF
		tokenEOF.Literal = ""

		tokenList := []Token.Token{}

		for {
			tok := l.NextToken()

			if tok == tokenEOF {
				break
			}
			tokenList = append(tokenList, tok)
		}

		assert.Equal(t, tc.want, tokenList, "The two words should be the same.")
	}
}