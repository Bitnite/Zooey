package token

type Type string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identificadores + Literais
	IDENT   = "IDENT"
	INT     = "INT"
	STRING  = "STRING"
	FLOAT   = "FLOAT"
	BOOLEAN = "BOOL"

	// Operators
	ASSIGN     = "ASSIGN"
	PLUS       = "+"
	PLUSPLUS   = "++"
	MINUSMINUS = "--"
	MINUS      = "-"
	BANG       = "!"
	ASTERISK   = "*"
	SLASH      = "/"

	LT  = "<"
	GT  = ">"
	LTE = "<="
	GTE = ">="
	POW = "^"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimitadores
	DOT       = "."
	SEMICOLON = ";"
	COLON     = ":"
	COMMA     = ","

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FN     = "FN"
	OwO    = "OwO"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	WHILE  = "WHILE"
	AND    = "&&"
	OR     = "||"
	FOR    = "FOR"
)

type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"owo":    OwO,
	"fn":     FN,
	"true":   TRUE,
	"false":  FALSE,
	"while":  WHILE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
}

func LookupIdent(identifier string) Type {
	if tok, ok := keywords[identifier]; ok {
		return tok
	} else {
		return IDENT
	}
}
