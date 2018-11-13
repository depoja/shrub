package repl

import (
	"bufio"
	"fmt"
	"os"
	"shrub/data"
	"shrub/evaluator"
	"shrub/parser"
	"strings"
)

func Printable(d data.Data) string {
	switch d := d.(type) {
	case []data.Data:
		l := make([]string, len(d))
		for i, x := range d {
			l[i] = Printable(x)
		}
		return "(" + strings.Join(l, " ") + ")"

	case data.Boolean:
		if d {
			return fmt.Sprint("#t")
		}
		return fmt.Sprint("#f")

	default:
		return fmt.Sprint(d)
	}
}

// Repl creates a repl
func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	environment := evaluator.NewEnvironment()

	for fmt.Print("> "); scanner.Scan(); fmt.Print("> ") {
		input := scanner.Text()
		program := parser.Parse(input)
		result := evaluator.Eval(program, environment)

		fmt.Println("=", Printable(result))
	}
}
