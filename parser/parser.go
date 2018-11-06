package parser

import (
	"shrub/data"
	"shrub/tokenizer"
	"strconv"
)

// Parse takes an input string, tokenizes it and parses the tokens
func Parse(input string) data.Data {
	tokens := tokenizer.Tokenize(input)
	return parseTokens(&tokens)
}

func parseTokens(tokens *[]string) data.Data {
	// Pop the first token
	token := (*tokens)[0]
	*tokens = (*tokens)[1:]

	switch token {
	// Denotes a list
	case "(":
		list := []data.Data{}
		for (*tokens)[0] != ")" {
			if i := parseTokens(tokens); i != data.Symbol("") {
				list = append(list, i)
			}
		}
		*tokens = (*tokens)[1:]
		return list

	// Denotes an atom
	default:
		if f, err := strconv.ParseFloat(token, 64); err == nil {
			return data.Number(f)
		}
		return data.Symbol(token)
	}
}
