package parser

import (
	"testing"

	"github.com/adamvinueza/monkey/ast"
	"github.com/adamvinueza/monkey/lexer"
)

func TestLetStatementGoodInput(t *testing.T) {
	input := `
let x = 5;
let y= 10;
let foobar = 838383;
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

func TestLetStatementBadInput(t *testing.T) {
	input := `
let x 5;
let = 10;
let 838383;
`
	p := New(lexer.New(input))

	expectedStatementsCount := 3
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != expectedStatementsCount {
		t.Fatalf("Expected program.Statements to have %d statements, found %d",
			expectedStatementsCount, len(program.Statements))
	}

	tests := []struct {
		expectedError string
	}{
		{"expected next token to be =, found INT"},
		{"expected next token to be IDENT, found ="},
		{"expected next token to be IDENT, found INT"},
	}

	errors := p.Errors()
	for i, tt := range tests {
		if errors[i] != tt.expectedError {
			t.Fatalf("Expected error msg '%s', found '%s'",
				tt.expectedError, errors[i])
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

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	expectedStatementsCount := 3
	p := New(lexer.New(input))
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != expectedStatementsCount {
		t.Fatalf("Expected program.Statements to have %d statements, found %d",
			expectedStatementsCount, len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, found %T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', found %q",
				returnStmt.TokenLiteral())
		}
	}
}
