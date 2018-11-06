package tokenizer

import "strings"

// Tokenize takes an input string and returns a slice of tokens
func Tokenize(input string) []string {
	// Adds whitespacing between parenthesis
	spaced := strings.Replace(input, "(", "( ", -1)
	spaced = strings.Replace(spaced, ")", " )", -1)

	// Split string into tokens
	return strings.Split(spaced, " ")
}
