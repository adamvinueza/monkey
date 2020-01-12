package token

import "fmt"

// TokenType represents the type of lexical token.
type TokenType string

// Token represents a lexical token, such as a keyword, operator, or semicolon.
type Token struct {
	Type    TokenType
	Literal string
}

// Returns the string value of this Token.
func (t Token) String() string {
	return fmt.Sprintf(`Type: [%s] | Literal: '%s'`, t.Type, t.Literal)
}

const (
	// Miscellaneous
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers & literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// Parentheses and Brackets
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent returns either the appropriate keyword TokenType or else IDENT.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
