package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
   `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got = %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func TestLetStatementsInvalidInput1(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let 838383;
   `

	l := lexer.New(input)
	p := New(l)

	p.ParseProgram()

	if len(p.Errors()) == 0 {
		t.Error("Parsed invalid input, yet no error detected")
	}
}

func TestLetStatementsInvalidInput2(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let . = ;
   `

	l := lexer.New(input)
	p := New(l)

	p.ParseProgram()

	if len(p.Errors()) == 0 {
		t.Error("Parsed invalid input, yet no error detected")
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 993322;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got = %d", len(program.Statements))
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

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiteral() is not 'let'. got = %q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement is not of *ast.LetStatement. got = %T", statement)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement name not '%s'. got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("statement not '%s'. got=%s", name, letStatement.Name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has encoutered %d errors", len(errors))
	for _, message := range errors {
		t.Errorf("parser error %s", message)
	}
	t.FailNow()
}
