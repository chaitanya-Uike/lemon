package ast

import "github.com/chaitanya-Uike/lemon/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statemets []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statemets) > 0 {
		return p.Statemets[0].TokenLiteral()
	}
	return ""
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
