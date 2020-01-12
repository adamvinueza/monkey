package parser

import (
	"testing"

	"github.com/adamvinueza/monkey/ast"
	"github.com/adamvinueza/monkey/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
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

	test := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatemnt(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
        t.Errorf("Expected s.TokenLiteral() to be 'let', found '%s'",
        s.TokenLteral())
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
        }
    }

}
