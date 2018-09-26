package workflowlang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokeniserWords(t *testing.T) {
	input := `show cancel alert`
	actual := NewTokeniser(nil).Stream(input)
	expected := []*Token{
		{col: 0, text: "show"}, {col: 4, whitespace: " "}, {col: 5, text: "cancel"}, {col: 11, whitespace: " "}, {col: 12, text: "alert"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserParens(t *testing.T) {
	input := `(Title) (Body)`
	actual := NewTokeniser([]byte{'(', ')'}).Stream(input)
	expected := []*Token{
		{col: 0, symbol: '('}, {col: 1, text: "Title"}, {col: 6, symbol: ')'}, {col: 7, whitespace: " "}, {col: 8, symbol: '('}, {col: 9, text: "Body"}, {col: 13, symbol: ')'},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserVariable(t *testing.T) {
	input := `$foo $bar`
	actual := NewTokeniser([]byte{'$'}).Stream(input)
	expected := []*Token{
		{col: 0, symbol: '$'}, {col: 1, text: "foo"}, {col: 4, whitespace: " "}, {col: 5, symbol: '$'}, {col: 6, text: "bar"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserEscapedChars(t *testing.T) {
	input := `start=\$bar something $foo`
	actual := NewTokeniser([]byte{'$', '\\'}).Stream(input)
	expected := []*Token{
		{col: 0, text: "start="}, {col: 6, symbol: '\\'}, {col: 7, symbol: '$'}, {col: 8, text: "bar"}, {col: 11, whitespace: " "}, {col: 12, text: "something"}, {col: 21, whitespace: " "}, {col: 22, symbol: '$'}, {col: 23, text: "foo"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserNewlines(t *testing.T) {
	input := `if $foo equals bar
	something
end if`
	actual := NewTokeniser([]byte{'$', '\\'}).Stream(input)
	expected := []*Token{
		{col: 0, text: "if"}, {col: 2, whitespace: " "}, {col: 3, symbol: '$'}, {col: 4, text: "foo"}, {col: 7, whitespace: " "}, {col: 8, text: "equals"}, {col: 14, whitespace: " "}, {col: 15, text: "bar"}, {col: 18, newline: true},
		{line: 1, col: 0, whitespace: "\t"}, {line: 1, col: 1, text: "something"}, {line: 1, col: 10, newline: true},
		{line: 2, col: 0, text: "end"}, {line: 2, col: 3, whitespace: " "}, {line: 2, col: 4, text: "if"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserTrailingNewline(t *testing.T) {
	input := `first
second
newlineafter
`
	actual := NewTokeniser(nil).Stream(input)
	expected := []*Token{
		{col: 0, text: "first"}, {col: 5, newline: true},
		{line: 1, col: 0, text: "second"}, {line: 1, col: 6, newline: true},
		{line: 2, col: 0, text: "newlineafter"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserWhitespace(t *testing.T) {
	input := "foo    bar wef"
	actual := NewTokeniser(nil).Stream(input)
	expected := []*Token{
		{col: 0, text: "foo"}, {col: 3, whitespace: "    "}, {col: 7, text: "bar"}, {col: 10, whitespace: " "}, {col: 11, text: "wef"},
	}
	assert.Equal(t, expected, actual)
}
