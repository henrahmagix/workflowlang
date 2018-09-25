package workflowlang

type Token struct {
	line int
	col  int
	text string
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
				stream = append(stream, &Token{line, pos, text})
				text = ""
			}
			pos = i - colReset
			stream = append(stream, &Token{line, pos, string(char)})
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
		stream = append(stream, &Token{line, pos, text})
	}
	return
}
