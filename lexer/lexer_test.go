package lexer

import (
	"testing"

	"github.com/hitochan777/monkey/token"
)

type ExpectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func compareToken(t *testing.T, testNum int, tok token.Token, expectedToken ExpectedToken) {
	t.Helper()

	if tok.Type != expectedToken.expectedType {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", testNum, expectedToken.expectedType, tok.Type)
	}

	if tok.Literal != expectedToken.expectedLiteral {
		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", testNum, expectedToken.expectedLiteral, tok.Literal)
	}

}

func TestNextToken(t *testing.T) {
	// test 1: symbols
	{
		input := `=+(){},;[]`

		tests := []ExpectedToken{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.LBRACKET, "["},
			{token.RBRACKET, "]"},
			{token.EOF, ""},
		}

		l := NewLexer(input)
		for i, tt := range tests {
			tok := l.NextToken()
			compareToken(t, i, tok, tt)
		}
	}
	// test 2
	{
		input := `let five = 5;
	let ten = 10;
	let add = fn(x,y) {
		x + y;
	};
	let result = add(five, ten);
	`

		tests := []ExpectedToken{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := NewLexer(input)
		for i, tt := range tests {
			tok := l.NextToken()
			compareToken(t, i, tok, tt)
		}
	}

	// test 3: operator test
	{
		input := `!-/*5;
		5 < 10 > 5;	
		`

		tests := []ExpectedToken{
			{token.BANG, "!"},
			{token.MINUS, "-"},
			{token.SLASH, "/"},
			{token.ASTERISK, "*"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.GT, ">"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := NewLexer(input)
		for i, tt := range tests {
			tok := l.NextToken()
			compareToken(t, i, tok, tt)
		}
	}

	// test 4: true, false, if, else, return
	{
		input := `if (5 < 10) {
			return true;	
		} else {
			return false;
		}
		`

		tests := []ExpectedToken{
			{token.IF, "if"},
			{token.LPAREN, "("},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.TRUE, "true"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.ELSE, "else"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.FALSE, "false"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.EOF, ""},
		}

		l := NewLexer(input)
		for i, tt := range tests {
			tok := l.NextToken()
			compareToken(t, i, tok, tt)
		}
	}

	// test 5: tests for double-character tokens
	{
		input := `10 == 10;
		10 != 9;
		`

		tests := []ExpectedToken{
			{token.INT, "10"},
			{token.EQ, "=="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.INT, "10"},
			{token.NOT_EQ, "!="},
			{token.INT, "9"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := NewLexer(input)
		for i, tt := range tests {
			tok := l.NextToken()
			compareToken(t, i, tok, tt)
		}
	}

	// test 6: tests for string ending with =
	{
		input := `10 =`

		tests := []ExpectedToken{
			{token.INT, "10"},
			{token.ASSIGN, "="},
		}

		l := NewLexer(input)
		for i, tt := range tests {
			tok := l.NextToken()
			compareToken(t, i, tok, tt)
		}
	}

	// test 7: test string
	{
		input := `"foobar"
		"foo bar"
		`

		tests := []ExpectedToken{
			{token.STRING, "foobar"},
			{token.STRING, "foo bar"},
			{token.EOF, ""},
		}

		l := NewLexer(input)
		for i, tt := range tests {
			tok := l.NextToken()
			compareToken(t, i, tok, tt)
		}

	}
}
