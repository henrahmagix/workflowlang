package workflowlang

import (
	"bufio"
	"bytes"
	"log"
	"strings"
	"unicode"
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

// var defaultSymbols = []byte{'(', ')', '$', '\\'}

func (t *tokeniser) Stream(input string) []*Token {
	var tokens []*Token

	allLines := []string{}
	r := bufio.NewScanner(strings.NewReader(input))
	for r.Scan() {
		allLines = append(allLines, r.Text())
	}
	if err := r.Err(); err != nil {
		log.Fatal(err)
	}

	lastline := len(allLines) - 1

	for line, linestr := range allLines {
		var (
			text       string
			whitespace string
			col        int
		)
		for i, char := range linestr {
			if bytes.ContainsRune(t.symbols, char) {
				if text != "" {
					tokens = append(tokens, &Token{line: line, col: col, text: text})
					text = ""
				}
				if whitespace != "" {
					tokens = append(tokens, &Token{line: line, col: col, whitespace: whitespace})
					whitespace = ""
				}
				col = i
				tokens = append(tokens, &Token{line: line, col: col, symbol: char})
				col += 1
			} else if unicode.IsSpace(char) {
				if text != "" {
					tokens = append(tokens, &Token{line: line, col: col, text: text})
					text = ""
					col = i
				}
				whitespace = whitespace + string(char)
			} else {
				if whitespace != "" {
					tokens = append(tokens, &Token{line: line, col: col, whitespace: whitespace})
					whitespace = ""
					col = i
				}
				text = text + string(char)
			}
		}
		if text != "" {
			tokens = append(tokens, &Token{line: line, col: col, text: text})
			text = ""
		}
		if whitespace != "" {
			tokens = append(tokens, &Token{line: line, col: col, whitespace: whitespace})
			whitespace = ""
		}
		if line != lastline {
			tokens = append(tokens, &Token{line: line, col: len(linestr), newline: true})
		}
	}

	return tokens
}
