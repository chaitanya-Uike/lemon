package parser

import (
	"testing"

	"github.com/chaitanya-Uike/lemon/ast"
	"github.com/chaitanya-Uike/lemon/lexer"
)

func TestReturnStatements(t *testing.T) {
	input := `
	return 5
	return 10
	return 993322`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statemets) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statemets))
	}

	for _, stmt := range program.Statemets {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
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
