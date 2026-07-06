package iteration

import (
	"strings"
)

const (
	Str    = "string"
	Wrt    = "write"
	WrtStr = "writeString"
)

func Repeat(text string, iterations int, method string) string {
	if iterations < 0 {
		iterations = 0
	}
	var repeated string

	if method == Wrt {
		var repeatedRaw strings.Builder
		for range iterations {
			repeatedRaw.Write([]byte(text))
		}
		repeated = repeatedRaw.String()
	} else if method == WrtStr {
		var repeatedRaw strings.Builder
		for range iterations {
			repeatedRaw.WriteString(text)
		}
		repeated = repeatedRaw.String()

	} else if method == Str {
		for range iterations {
			repeated += text
		}
	} else {
		return ""
	}
	return repeated
}

// TODO: Implement WriteString and string (for benchmark purposes)
