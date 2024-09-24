func removeStars(s string) string {
	var stack []byte
	// Loop around to find * and POP that * and closest element to star
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

