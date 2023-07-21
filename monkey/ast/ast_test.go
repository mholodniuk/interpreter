package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "var1"},
					Value: "var1",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "var2"},
					Value: "var2",
				},
			},
		},
	}

	if program.String() != "let var1 = var2;" {
		t.Errorf("program.String() is not correct, got = %q", program.String())
	}
}
