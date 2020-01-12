package parser

import (
	"fmt"
	"testing"

	"github.com/adamvinueza/monkey/ast"
	"github.com/adamvinueza/monkey/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
let x 5;
let = 10;
let 838383;
`
	p := New(lexer.New(input))

	expectedStatementsCount := 3
	program := p.ParseProgram()
    checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != expectedStatementsCount {
		t.Fatalf("Expected program.Statements to have %d statements, found %d",
			expectedStatementsCount, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors()
    if len(errors) == 0 {
        return
    }

    t.Errorf("parser has %d errors", len(errors))
    for _, msg := range errors {
        t.Errorf("parser error: %q", msg)
    }
    t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	l := s.(*ast.LetStatement)
	fmt.Printf("testLetStatement for '%s', name='%s', value='%s'\n",
		l.TokenLiteral(), l.Name, l.Value)
	if s.TokenLiteral() != "let" {
		t.Errorf("Expected s.TokenLiteral() to be 'let', found '%s'",
			s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.Statement, found %T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', found '%s'", name,
			letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s', found '%s'", name,
			letStmt.Name.TokenLiteral())
		return false
	}
	return true
}
