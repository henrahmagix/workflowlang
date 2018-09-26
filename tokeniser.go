package workflowlang

import (
	"bufio"
	"bytes"
	"log"
	"strings"
)

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

// var defaultSymbols = []byte{'(', ')', '$', '\\'}

func (t *tokeniser) Stream(input string) (stream []*Token) {
	text := ""
	space := ""
	col := 0
	line := -1

	all := []string{}
	r := bufio.NewScanner(strings.NewReader(input))
	for r.Scan() {
		all = append(all, r.Text())
	}
	if err := r.Err(); err != nil {
		log.Fatal(err)
	}

	lastline := len(all) - 1
	for line, linestr := range all {
		col = 0
		for i, char := range linestr {
			if bytes.ContainsRune(t.symbols, char) {
				if text != "" {
					stream = append(stream, &Token{line: line, col: col, text: text})
					text = ""
				}
				if space != "" {
					stream = append(stream, &Token{line: line, col: col, whitespace: space})
					space = ""
				}
				col = i
				stream = append(stream, &Token{line: line, col: col, symbol: char})
				col += 1
			} else if bytes.ContainsRune(whitespace, char) {
				if text != "" {
					stream = append(stream, &Token{line: line, col: col, text: text})
					text = ""
					col = i
				}
				space = space + string(char)
			} else {
				if space != "" {
					stream = append(stream, &Token{line: line, col: col, whitespace: space})
					space = ""
					col = i
				}
				text = text + string(char)
			}
		}
		if text != "" {
			stream = append(stream, &Token{line: line, col: col, text: text})
			text = ""
		}
		if space != "" {
			stream = append(stream, &Token{line: line, col: col, whitespace: space})
			space = ""
		}
		if line != lastline {
			stream = append(stream, &Token{line: line, col: len(linestr), newline: true})
		}
	}

	if text != "" {
		stream = append(stream, &Token{line: line, col: col, text: text})
	}
	if space != "" {
		stream = append(stream, &Token{line: line, col: col, whitespace: space})
	}
	return
}
