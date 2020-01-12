package ast

import (
	"bytes"
	"github.com/adamvinueza/monkey/token"
)

// Node is the fundamental unit in an abstract syntax tree.
type Node interface {
	// TokenLiteral returns the literal representing the lexical token that
	// represents this Node.
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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

// String returns the string value of this LetStatement.
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement represents the structure of a return statement, such as
// "return x + y;"
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement represents the structure of an expression, considered as
// a statement. (Expressions are also Statements in Monkey.)
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier represents the name of a variable.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// IntegerLiteral represents an integer literal, such as "5" or "65536".
type IntegerLiteral struct {
	Token token.Token // the token.INT token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
