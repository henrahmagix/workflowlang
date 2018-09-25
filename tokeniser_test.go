package workflowlang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokeniserWords(t *testing.T) {
	input := `show cancel alert`
	actual := TokenStream(input)
	expected := []*Token{
		{0, "show"}, {5, "cancel"}, {12, "alert"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserParens(t *testing.T) {
	input := `(Title) (Body)`
	actual := TokenStream(input)
	expected := []*Token{
		{0, "("}, {1, "Title"}, {6, ")"}, {8, "("}, {9, "Body"}, {13, ")"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserVariable(t *testing.T) {
	input := `$foo $bar`
	actual := TokenStream(input)
	expected := []*Token{
		{0, "$"}, {1, "foo"}, {5, "$"}, {6, "bar"},
	}
	assert.Equal(t, expected, actual)
}
