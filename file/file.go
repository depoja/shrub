package file

import (
	"bufio"
	"fmt"
	"os"
	"shrub/data"
	"shrub/evaluator"
	"shrub/parser"
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

	fmt.Println("=", printable(result))
}

func printable(d data.Data) string {
	switch d := d.(type) {
	case []data.Data:
		l := make([]string, len(d))
		for i, x := range d {
			l[i] = printable(x)
		}
		return "(" + strings.Join(l, " ") + ")"

	default:
		return fmt.Sprint(d)
	}
}
