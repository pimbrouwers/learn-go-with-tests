package iteration

func Repeat(char string, iter int) string {
	var repeated string

	for i := 0; i < iter; i++ {
		repeated += char
	}

	return repeated
}
