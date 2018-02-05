package evaluator

import (
	"testing"

	"github.com/hitochan777/monkey/lexer"
	"github.com/hitochan777/monkey/object"
	"github.com/hitochan777/monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testObject(t *testing.T, obj object.Object, expected interface{}) bool {
	switch v := expected.(type) {
	case int64:
		return testIntegerObject(t, obj, v)
	default:
		t.Errorf("Unexpected type %T", v)
		return false
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}
