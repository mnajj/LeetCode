func interpret(command string) string {
	var buf bytes.Buffer
	for i := 0; i < len(command); {
		switch r := command[i]; r {
		case 'G':
			buf.WriteRune('G')
			i++
		case '(':
			if rune(command[i+1]) == 'a' {
				buf.WriteString("al")
				i += 4
			} else {
				buf.WriteRune('o')
				i += 2
			}
		default:
			panic(r)
		}
	}
	return buf.String()
}
