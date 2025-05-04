package iteration

func Repeat(char string) string {
	var chars string
	for i := 0; i < 5; i++ {
		chars = chars + char
	}

	return chars
}
