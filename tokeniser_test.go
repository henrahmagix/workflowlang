package workflowlang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokeniserWords(t *testing.T) {
	input := `show cancel alert`
	actual := TokenStream(input)
	expected := []*Token{
		{col: 0, text: "show"}, {col: 4, text: " "}, {col: 5, text: "cancel"}, {col: 11, text: " "}, {col: 12, text: "alert"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserParens(t *testing.T) {
	input := `(Title) (Body)`
	actual := TokenStream(input)
	expected := []*Token{
		{col: 0, text: "("}, {col: 1, text: "Title"}, {col: 6, text: ")"}, {col: 7, text: " "}, {col: 8, text: "("}, {col: 9, text: "Body"}, {col: 13, text: ")"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserVariable(t *testing.T) {
	input := `$foo $bar`
	actual := TokenStream(input)
	expected := []*Token{
		{col: 0, text: "$"}, {col: 1, text: "foo"}, {col: 4, text: " "}, {col: 5, text: "$"}, {col: 6, text: "bar"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserEscapedChars(t *testing.T) {
	input := `start=\$bar something $foo`
	actual := TokenStream(input)
	expected := []*Token{
		{col: 0, text: "start="}, {col: 6, text: "\\"}, {col: 7, text: "$"}, {col: 8, text: "bar"}, {col: 11, text: " "}, {col: 12, text: "something"}, {col: 21, text: " "}, {col: 22, text: "$"}, {col: 23, text: "foo"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserNewlines(t *testing.T) {
	input := `if $foo equals bar
	something
end if`
	actual := TokenStream(input)
	expected := []*Token{
		{line: 0, col: 0, text: "if"}, {line: 0, col: 2, text: " "}, {line: 0, col: 3, text: "$"}, {line: 0, col: 4, text: "foo"}, {line: 0, col: 7, text: " "}, {line: 0, col: 8, text: "equals"}, {line: 0, col: 14, text: " "}, {line: 0, col: 15, text: "bar"}, {line: 0, col: 18, text: "\n"},
		{line: 1, col: 0, text: "\t"}, {line: 1, col: 1, text: "something"}, {line: 1, col: 10, text: "\n"},
		{line: 2, col: 0, text: "end"}, {line: 2, col: 3, text: " "}, {line: 2, col: 4, text: "if"},
	}
	assert.Equal(t, expected, actual)
}
