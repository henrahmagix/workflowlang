package workflowlang

import "bytes"

type Token struct {
	line       int
	col        int
	symbol     rune
	text       string
	whitespace string
	newline    bool
}

type tokeniser struct {
	symbols []byte
}

func NewTokeniser(symbols []byte) *tokeniser {
	return &tokeniser{symbols}
}

var whitespace = []byte{' ', '\t'}
var newlines = []byte{'\r', '\n'}

// var defaultSymbols = []byte{'(', ')', '$', '\\', '\n', '\t'}

func (t *tokeniser) Stream(input string) (stream []*Token) {
	text := ""
	space := ""
	col := 0
	line := 0
	colReset := 0
	for i, char := range input {
		if bytes.ContainsRune(t.symbols, char) {
			if text != "" {
				stream = append(stream, &Token{line: line, col: col, text: text})
				text = ""
			}
			if space != "" {
				stream = append(stream, &Token{line: line, col: col, whitespace: space})
				space = ""
			}
			col = i - colReset
			stream = append(stream, &Token{line: line, col: col, symbol: char})
			col += 1
		} else if bytes.ContainsRune(newlines, char) {
			if text != "" {
				stream = append(stream, &Token{line: line, col: col, text: text})
				text = ""
			}
			if space != "" {
				stream = append(stream, &Token{line: line, col: col, whitespace: space})
				space = ""
			}
			col = i - colReset
			stream = append(stream, &Token{line: line, col: col, newline: true})
			col = 0
			colReset = i + 1
			line += 1
		} else if bytes.ContainsRune(whitespace, char) {
			if text != "" {
				stream = append(stream, &Token{line: line, col: col, text: text})
				text = ""
				col = i - colReset
			}
			space = space + string(char)
		} else {
			if space != "" {
				stream = append(stream, &Token{line: line, col: col, whitespace: space})
				space = ""
				col = i - colReset
			}
			text = text + string(char)
		}
	}
	if text != "" {
		stream = append(stream, &Token{line: line, col: col, text: text})
	}
	return
}
