package iteration

func Repeat(character string, amount int) string {
	var repeated string
	for i := 0; i < amount; i++ {
		repeated += character
	}

	return repeated
}
