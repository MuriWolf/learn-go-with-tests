package iteration

func Repeat(text string) string {
	var repeated string
	for range 5 {
		repeated += text
	}
	return repeated
}
