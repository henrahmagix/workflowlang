package workflowlang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestTokeniser(t *testing.T) {
// 	input := `if $foo contains(some string)
// 	show cancel alert (Title message) (This is the body showing variable \$bar as $bar with text after)
// end if
// show $$date(current)`
// 	stream := tokens(input)
// 	expected := []string{"if", "$", "foo", "contains", "(", "some string", ")",
// 		"show", "cancel", "alert", "(", "Title message", ")",
// 		"(", "This is the body showing variable $bar as ", "$", "bar", " with text after", ")",
// 		"end", "if",
// 		"show", "$$", "date(current)",
// 	}
// }

func TestTokeniserWords(t *testing.T) {
	input := `show cancel alert`
	actual := TokenStream(input)
	expected := []*Token{
		{0, "show"}, {4, " "}, {5, "cancel"}, {11, " "}, {12, "alert"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserParens(t *testing.T) {
	input := `(Title) (Body)`
	actual := TokenStream(input)
	expected := []*Token{
		{0, "("}, {1, "Title"}, {6, ")"}, {7, " "}, {8, "("}, {9, "Body"}, {13, ")"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserVariable(t *testing.T) {
	input := `$foo $bar`
	actual := TokenStream(input)
	expected := []*Token{
		{0, "$"}, {1, "foo"}, {4, " "}, {5, "$"}, {6, "bar"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserEscapedChars(t *testing.T) {
	input := `start=\$bar something $foo`
	actual := TokenStream(input)
	expected := []*Token{
		{0, "start="}, {6, "\\"}, {7, "$"}, {8, "bar"}, {11, " "}, {12, "something"}, {21, " "}, {22, "$"}, {23, "foo"},
	}
	assert.Equal(t, expected, actual)
}

func TestTokeniserNewlines(t *testing.T) {
	input := `if $foo equals bar
	something
end if`
	actual := TokenStream(input)
	expected := []*Token{
		{0, "if"}, {2, " "}, {3, "$"}, {4, "foo"}, {7, " "}, {8, "equals"}, {14, " "}, {15, "bar"}, {18, "\n"},
		{0, "\t"}, {1, "something"}, {10, "\n"},
		{0, "end"}, {3, " "}, {4, "if"},
	}
	assert.Equal(t, expected, actual)
}
