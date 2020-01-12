package parser

import (
	"github.com/adamvinueza/ast"
	"github.com/adamvinueza/lexer"
	"github.com/adamvinueza/token"
)

type Parser struct {
	l *lexer.lexer

	curToken  token.Token
	peekToken token.Token

    func New(l *lexer.Lexer) *Parser {
        p := &Parser(l: l}

        p.nextToken()
        p.nextToken()

        return p
    }

    func (p *Parser) nextToken() {
        p.curToken = p.peekToken
        p.peekToken = p.l.NextToken()
    }

    func (p *Parser) ParseProgram() *ast.Program {
        return nil
    }
}
