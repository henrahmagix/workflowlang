package workflowlang

type Token struct {
	position int
	text     string
}

func TokenStream(input string) (stream []*Token) {
	pos := 0
	text := ""
	lineReset := 0
	for i, char := range input {
		switch char {
		case ' ', '(', ')', '$', '\\', '\n', '\t':
			if text != "" {
				stream = append(stream, &Token{pos, text})
				text = ""
			}
			pos = i - lineReset
			stream = append(stream, &Token{pos, string(char)})
			if char == '\n' {
				pos = 0
				lineReset = i + 1
			} else {
				pos = pos + 1
			}
		default:
			text = text + string(char)
		}
	}
	if text != "" {
		stream = append(stream, &Token{pos, text})
	}
	return
}
