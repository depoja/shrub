package file

import (
	"bufio"
	"fmt"
	"os"
	"shrub/data"
	"shrub/evaluator"
	"shrub/parser"
	"shrub/repl"
	"strings"
)

// Interpret reads a file with the given fileName and returns the programs result
func Interpret(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	environment := evaluator.NewEnvironment()
	var result data.Data

	for scanner.Scan() {
		input := scanner.Text()
		if !strings.HasPrefix(input, ";") && len(input) > 0 {
			program := parser.Parse(string(input))
			result = evaluator.Eval(program, environment)
		}
	}

	fmt.Println("=", repl.Printable(result))
}
