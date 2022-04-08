package parser

import (
	"monkey/ast"
	"monkey/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	return nil
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	return nil
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return nil
}
