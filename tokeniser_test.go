package workflowlang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokeniser(t *testing.T) {
	input := `show cancel alert (Title) (Body)`
	actual := TokenStream(input)
	expected := []*Token{
		{0, "show"}, {5, "cancel"}, {12, "alert"}, {18, "("}, {19, "Title"}, {24, ")"}, {26, "("}, {27, "Body"}, {31, ")"},
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
