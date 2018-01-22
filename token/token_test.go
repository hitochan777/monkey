package token

import "testing"

func TestLookupIdent(t *testing.T) {
	// test 1: returns IDENT when the input is not a keyword
	{
		input := "hello"
		var expectedTokenType TokenType = IDENT
		if tokenType := LookupIdent(input); tokenType != expectedTokenType {
			t.Fatalf("tokentype wrong. expected=%q, got=%q", expectedTokenType, tokenType)
		}

	}

	// test 2: returns FUNCTION when the input is fn
	{
		input := "fn"
		var expectedTokenType TokenType = FUNCTION
		if tokenType := LookupIdent(input); tokenType != expectedTokenType {
			t.Fatalf("tokentype wrong. expected=%q, got=%q", expectedTokenType, tokenType)
		}

	}

}
