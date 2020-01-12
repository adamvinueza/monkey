package ast

import "github.com/adamvinueza/monkey/token"

// Node is the fundamental unit in an abstract syntax tree.
type Node interface {
	// TokenLiteral returns the literal representing the lexical token that
	// represents this Node.
	TokenLiteral() string
}

// Statement is a program statement, such as "let x = 5;".
type Statement interface {
	Node
	statementNode()
}

// Expression is an expression--it returns a value. An example is "x + y".
type Expression interface {
	Node
	expressionNode()
}

// Program represents a Monkey program. It is produced by a parser.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the TokenLiteral of the first Statement, or the empty
// string if there are no Statements.
//
// For example, any implementation of a let statement should have the token
// literal "let", so if the first statement in the program is a let statement,
// TokenLiteral should return "let".
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement represents the structure of a let statement. It is of the form
// LET NAME = VALUE.
//
// For example, in "let a = b", "a" represents the identifier and "b" represents
// the expression.
//
// Appropriately, LetStatement.TokenLiteral returns "LET".
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier represents the name of a variable.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) statementNode()       {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
