package Token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identificadores + Literais
	IDENT   = "IDENT" // int -> x <- = 5
	INT     = "INT"   // 123456789 ...
	STRING  = "STRING"
	FLOAT   = "FLOAT" // Implementar dps
	BOOLEAN = "BOOL"

	// Operators
	ASSIGN   = "ASSIGN"
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT  = "<"
	GT  = ">"
	LTE = "<="
	GTE = ">="

	EQ     = "=="
	NOT_EQ = "!="

	// Delimitadores
	DOT       = "." // implementar dps
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FN     = "FN"  //Function
	OwO    = "OwO" //let -> pensar em nome melhor
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	WHILE  = "WHILE"
	AND    = "&&"
	OR     = "||"
	FOR    = "FOR" //implementar dps
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
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

func LookupIdent(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	} else {
		return IDENT
	}
}
