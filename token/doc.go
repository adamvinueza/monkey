// Package token holds the lexical token types for use in lexing and parsing,
// and functions for accessing them.
//
// Sample token types include IDENT for an identifier, TRUE for the keyword
// "true", and FUNCTION for the keyword "fn".
//
// A Token represents a particular lexical token. It has a type and a literal:
// for example, Token{Type: IF, Literal: "if"} represents a token identified as
// an IF token.
package token
