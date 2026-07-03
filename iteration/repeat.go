package iteration

import "strings"

func Repeat(text string, iterations int) string {
	if iterations < 0 {
		iterations = 0
	}

	var repeated strings.Builder
	for range iterations {
		repeated.Write([]byte(text))
	}
	return repeated.String()
}
