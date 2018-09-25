package workflowlang

import "bytes"

type Token struct {
	line       int
	col        int
	symbol     rune
	text       string
	whitespace int
}

type tokeniser struct {
	symbols []byte
}

func NewTokeniser(symbols []byte) *tokeniser {
	return &tokeniser{symbols}
}

var whitespace = []byte{' ', '\n', '\t'}

// var defaultSymbols = []byte{'(', ')', '$', '\\', '\n', '\t'}

func (t *tokeniser) stream(input string) (stream []*Token) {
	pos := 0
	text := ""
	line := 0
	colReset := 0
	for i, char := range input {
		if bytes.ContainsRune(t.symbols, char) || bytes.ContainsRune(whitespace, char) {
			if text != "" {
				stream = append(stream, &Token{line: line, col: pos, text: text})
				text = ""
			}
			pos = i - colReset
			t := &Token{line: line, col: pos}
			if char == ' ' {
				t.whitespace += 1
			} else {
				t.symbol = char
			}
			stream = append(stream, t)
			if char == '\n' {
				pos = 0
				colReset = i + 1
				line += 1
			} else {
				pos += 1
			}
		} else {
			text = text + string(char)
		}
	}
	if text != "" {
		stream = append(stream, &Token{line: line, col: pos, text: text})
	}
	return
}
