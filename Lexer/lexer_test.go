package Lexer

import (
	"testing"

	token "github.com/ZooeyLang/Token"
	"github.com/stretchr/testify/assert"
)

func TestLexer_NextToken(t *testing.T) {
	type test struct {
		name    string
		input   string
		want    []token.Token
		wantErr bool
	}

	tests := []test{
		{
			name:  "should tokenize simple expression correctly",
			input: "owo teste :=: 5;",
			want: []token.Token{
				{Type: token.OwO, Literal: "owo"},
				{Type: token.IDENT, Literal: "teste"},
				{Type: token.ASSIGN, Literal: ":=:"},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},
			},
			wantErr: false,
		},
		{
			name:  "should tokenize if GT expression correctly",
			input: "if 5 > 2 { return true }",
			want: []token.Token{
				{Type: token.IF, Literal: "if"},
				{Type: token.INT, Literal: "5"},
				{Type: token.GT, Literal: ">"},
				{Type: token.INT, Literal: "2"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RETURN, Literal: "return"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.RBRACE, Literal: "}"},
			},
			wantErr: false,
		},
		{
			name:  "should tokenize if EQ expression correctly",
			input: "if 2 == 2 { return true }",
			want: []token.Token{
				{Type: token.IF, Literal: "if"},
				{Type: token.INT, Literal: "2"},
				{Type: token.EQ, Literal: "=="},
				{Type: token.INT, Literal: "2"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RETURN, Literal: "return"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.RBRACE, Literal: "}"},
			},
			wantErr: false,
		},
		{
			name:  "should tokenize if NOT_EQ expression correctly",
			input: "if 5 != 2 { return true }",
			want: []token.Token{
				{Type: token.IF, Literal: "if"},
				{Type: token.INT, Literal: "5"},
				{Type: token.NOT_EQ, Literal: "!="},
				{Type: token.INT, Literal: "2"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RETURN, Literal: "return"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.RBRACE, Literal: "}"},
			},
			wantErr: false,
		},
		{
			name:  "should tokenize float value correctly",
			input: "owo myFloat :=: 5.25",
			want: []token.Token{
				{Type: token.OwO, Literal: "owo"},
				{Type: token.IDENT, Literal: "myFloat"},
				{Type: token.ASSIGN, Literal: ":=:"},
				{Type: token.FLOAT, Literal: "5.25"},
			},
			wantErr: false,
		},
		{
			name:  "should tokenize integer values correctly",
			input: "owo myInt :=: 5",
			want: []token.Token{
				{Type: token.OwO, Literal: "owo"},
				{Type: token.IDENT, Literal: "myInt"},
				{Type: token.ASSIGN, Literal: ":=:"},
				{Type: token.INT, Literal: "5"},
			},
			wantErr: false,
		},
		{
			name:  "should tokenize all tokens correctly",
			input: `owo x 5 "xx" 10.25 true false :=: + - ! * / > < >= <= == != ; : fn && for while`,
			want: []token.Token{
				{Type: token.OwO, Literal: "owo"},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.INT, Literal: "5"},
				{Type: token.STRING, Literal: "xx"},
				{Type: token.FLOAT, Literal: "10.25"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.FALSE, Literal: "false"},
				{Type: token.ASSIGN, Literal: ":=:"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.BANG, Literal: "!"},
				{Type: token.ASTERISK, Literal: "*"},
				{Type: token.SLASH, Literal: "/"},
				{Type: token.GT, Literal: ">"},
				{Type: token.LT, Literal: "<"},
				{Type: token.GTE, Literal: ">="},
				{Type: token.LTE, Literal: "<="},
				{Type: token.EQ, Literal: "=="},
				{Type: token.NOT_EQ, Literal: "!="},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.FN, Literal: "fn"},
				{Type: token.AND, Literal: "&&"},
				{Type: token.FOR, Literal: "for"},
				{Type: token.WHILE, Literal: "while"},
			},
			wantErr: false,
		},
		{
			name:  "inexistent token should be illegal",
			input: ":=",
			want: []token.Token{
				{Type: token.ILLEGAL, Literal: "="},
			},
			wantErr: true,
		},

		{
			name:  "should tokenize dot and comma",
			input: ". ,",
			want: []token.Token{
				{Type: token.DOT, Literal: "."},
				{Type: token.COMMA, Literal: ","},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.input)

			var tokenEOF token.Token
			tokenEOF.Type = token.EOF
			tokenEOF.Literal = ""

			tokenList := []token.Token{}

			for {
				tok := l.NextToken()

				if tok == tokenEOF {
					break
				}
				tokenList = append(tokenList, tok)
			}

			assert.Equal(t, tc.want, tokenList, "The tokenList must be equal!")
		})
	}
}
