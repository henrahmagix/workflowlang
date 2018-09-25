package workflowlang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokeniserWords(t *testing.T) {
	input := `show cancel alert`
	actual := NewTokeniser(nil).stream(input)
	expected := []*Token{
		{col: 0, text: "show"}, {col: 4, text: " "}, {col: 5, text: "cancel"}, {col: 11, text: " "}, {col: 12, text: "alert"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserParens(t *testing.T) {
	input := `(Title) (Body)`
	actual := NewTokeniser([]byte{'(', ')'}).stream(input)
	expected := []*Token{
		{col: 0, symbol: '('}, {col: 1, text: "Title"}, {col: 6, symbol: ')'}, {col: 7, text: " "}, {col: 8, symbol: '('}, {col: 9, text: "Body"}, {col: 13, symbol: ')'},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserVariable(t *testing.T) {
	input := `$foo $bar`
	actual := NewTokeniser([]byte{'$'}).stream(input)
	expected := []*Token{
		{col: 0, symbol: '$'}, {col: 1, text: "foo"}, {col: 4, text: " "}, {col: 5, symbol: '$'}, {col: 6, text: "bar"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserEscapedChars(t *testing.T) {
	input := `start=\$bar something $foo`
	actual := NewTokeniser([]byte{'$', '\\'}).stream(input)
	expected := []*Token{
		{col: 0, text: "start="}, {col: 6, symbol: '\\'}, {col: 7, symbol: '$'}, {col: 8, text: "bar"}, {col: 11, text: " "}, {col: 12, text: "something"}, {col: 21, text: " "}, {col: 22, symbol: '$'}, {col: 23, text: "foo"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserNewlines(t *testing.T) {
	input := `if $foo equals bar
	something
end if`
	actual := NewTokeniser([]byte{'$', '\\'}).stream(input)
	expected := []*Token{
		{col: 0, text: "if"}, {col: 2, text: " "}, {col: 3, symbol: '$'}, {col: 4, text: "foo"}, {col: 7, text: " "}, {col: 8, text: "equals"}, {col: 14, text: " "}, {col: 15, text: "bar"}, {col: 18, symbol: '\n'},
		{line: 1, col: 0, symbol: '\t'}, {line: 1, col: 1, text: "something"}, {line: 1, col: 10, symbol: '\n'},
		{line: 2, col: 0, text: "end"}, {line: 2, col: 3, text: " "}, {line: 2, col: 4, text: "if"},
	}
	assert.Equal(t, expected, actual)
}
