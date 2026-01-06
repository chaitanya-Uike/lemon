package parser

import (
	"fmt"
	"testing"

	"github.com/chaitanya-Uike/lemon/ast"
	"github.com/chaitanya-Uike/lemon/lexer"
)

func TestIdentifierLiteral(t *testing.T) {
	input := `myVar`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Expected %d statements, got %d", 1, len(program.Statements))
	}

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected *ast.ExpressionStatement, got %T", exp)
	}

	testIdentiferLiteral(t, exp.Expression, "myVar")
}

func testIdentiferLiteral(t *testing.T, exp ast.Expression, value string) {
	ident, ok := exp.(*ast.IdentifierLiteral)
	if !ok {
		t.Fatalf("Expected *ast.IdentifierLiteral, got %T", ident)
	}

	if ident.Value != value {
		t.Fatalf("ident.Value not %s. got %s", value, ident.Value)
	}

	if ident.TokenLiteral() != value {
		t.Fatalf("ident.TokenLiteral not %s. got %s", value, ident.TokenLiteral())
	}
}

func TestIntegerLiteral(t *testing.T) {
	input := `5`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Expected %d statements, got %d", 1, len(program.Statements))
	}

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected *ast.ExpressionStatement, go %T", exp)
	}

	testIntegerLiteral(t, exp.Expression, 5)

}

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) {
	intLiteral, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("Expected *ast.IntegerLiteral, got %T", intLiteral)
	}

	if intLiteral.Value != value {
		t.Errorf("literal.Value not %d. got=%d", value, intLiteral.Value)
	}

	if intLiteral.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("literal.TokenLiteral not %d. got=%s", value,
			intLiteral.TokenLiteral())
	}
}

func TestFloatLiteral(t *testing.T) {
	input := `5.25`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Expected %d statements, got %d", 1, len(program.Statements))
	}

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected *ast.ExpressionStatement, got %T", exp)
	}

	testFloatLiteral(t, exp.Expression, 5.25)

}

func testFloatLiteral(t *testing.T, exp ast.Expression, value float64) {
	floatLiteral, ok := exp.(*ast.FloatLiteral)
	if !ok {
		t.Fatalf("Expected *ast.FloatLiteral, got %T", floatLiteral)
	}

	if floatLiteral.Value != value {
		t.Errorf("literal.Value not %f. got=%f", value, floatLiteral.Value)
	}

	if floatLiteral.TokenLiteral() != fmt.Sprintf("%g", value) {
		t.Errorf("literal.TokenLiteral not %g. got=%s", value,
			floatLiteral.TokenLiteral())
	}
}

func TestBooleanLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("Expeccted %d statements, got %d", 1, len(program.Statements))
		}

		exp, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Expected *ast.ExpressionStatement, got %T", exp)
		}

		testBooleanLiteral(t, exp.Expression, tt.expected)
	}

}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) {
	boolLiteral, ok := exp.(*ast.BooleanLiteral)
	if !ok {
		t.Fatalf("Expected *ast.BooleanLiteral, got %T", boolLiteral)
	}

	if boolLiteral.Value != value {
		t.Fatalf("Expected boolLiteral.Value=%t, got %t", value, boolLiteral.Value)
	}

	if boolLiteral.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Fatalf("Expected boolLiteral.TokenLiteral=%t, got %s", value, boolLiteral.TokenLiteral())
	}
}

func TestPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    any
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
		{"!true", "!", true},
		{"!false", "!", false},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("Expected %d statements. got=%d",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Expected *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Expected ast.PrefixExpression. got=%T", stmt.Expression)
		}

		testLiteralExpression(t, exp.Expression, tt.value)
	}
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected any) {
	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, exp, int64(v))
	case int64:
		testIntegerLiteral(t, exp, v)
	case float32:
		testFloatLiteral(t, exp, float64(v))
	case float64:
		testFloatLiteral(t, exp, v)
	case bool:
		testBooleanLiteral(t, exp, v)
	case string:
		testIdentiferLiteral(t, exp, v)
	default:
		t.Fatalf("Type %T not handled", exp)
	}
}

func TestInfixExpression(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  any
		operator   string
		rightValue any
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("Expected %d statements. got=%d",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Expected *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue)
	}
}

func testInfixExpression(t *testing.T, exp ast.Expression, left any, operator string, right any) {
	inExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Expected *ast.InfixExpression. got=%T", inExp)
	}

	if inExp.Operator != operator {
		t.Fatalf("Expected operator '%s'. got=%s",
			operator, inExp.Operator)
	}

	testLiteralExpression(t, inExp.Left, left)
	testLiteralExpression(t, inExp.Right, right)
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4\n -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5
	return 10
	return 993322`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
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
