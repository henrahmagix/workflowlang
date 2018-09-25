package workflowlang

type Token struct {
	line   int
	col    int
	symbol rune
	text   string
}

func TokenStream(input string) (stream []*Token) {
	pos := 0
	text := ""
	line := 0
	colReset := 0
	for i, char := range input {
		switch char {
		case ' ', '(', ')', '$', '\\', '\n', '\t':
			if text != "" {
				stream = append(stream, &Token{line: line, col: pos, text: text})
				text = ""
			}
			pos = i - colReset
			t := &Token{line: line, col: pos}
			if char == ' ' {
				t.text = string(char)
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
		default:
			text = text + string(char)
		}
	}
	if text != "" {
		stream = append(stream, &Token{line: line, col: pos, text: text})
	}
	return
}
