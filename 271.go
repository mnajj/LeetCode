func encode(strs []string) string {
	var buf bytes.Buffer
	for _, s := range strs {
		buf.WriteString(
			fmt.Sprintf("%d#%s", len(s), s))
	}
	return buf.String()
}

func decode(s string) []string {
	var result []string
	var num bytes.Buffer

	for i := 0; i < len(s); {
		if i >= len(s) {
			break
		}
		r := rune(s[i])
		if r != '#' {
			num.WriteRune(r)
			i++
			continue
		}
		i++
		ct, err := strconv.Atoi(num.String())
		if err != nil {
			panic(err)
		}
		result = append(result, s[i:i+ct])
		num.Reset()
		i += ct
	}
	return result
}
