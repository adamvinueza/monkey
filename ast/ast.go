package ast

import "github.com/adamvinueza/monkey/token"

// Node is the fundamental unit in our ASTs.
type Node interface {
	TokenLiteral() string
}

// Statement is a statement.
type Statement interface {
	Node
	statementNode()
}

// Expression is an expression--it returns a value.
type Expression interface {
	Node
	expressionNode()
}

// Program represents a Monkey program.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement represents the structure of a let statement.
//
// For example, in "let a = b", "a" represents the identifier and "b" represents
// the expression.
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
