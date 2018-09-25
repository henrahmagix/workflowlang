package workflowlang

type Token struct {
	position int
	text     string
}

func TokenStream(input string) (stream []*Token) {
	pos := 0
	text := ""
	for i, char := range input {
		switch char {
		case ' ':
			if text != "" {
				stream = append(stream, &Token{pos, text})
				text = ""
			}
			pos = i + 1
			continue
		case '(', ')', '$', '\\':
			if text != "" {
				stream = append(stream, &Token{pos, text})
				text = ""
			}
			stream = append(stream, &Token{i, string(char)})
			pos = i + 1
			continue
		}
		text = text + string(char)
	}
	if text != "" {
		stream = append(stream, &Token{pos, text})
	}
	return
}
