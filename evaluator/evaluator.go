package evaluator

import (
	"log"
	"shrub/data"
)

// Eval evaluates a given expression with a given environment
func Eval(expression data.Data, env *Environment) data.Data {
	var result data.Data

	switch e := expression.(type) {
	case data.Number:
		result = e

	case data.Symbol:
		result = env.Get(e).vars[e]

	case []data.Data:
		switch car, _ := e[0].(data.Symbol); car {
		case "quote":
			result = e[1]

		case "if":
			if Eval(e[1], env).(bool) {
				result = Eval(e[2], env)
			} else {
				result = Eval(e[3], env)
			}

		case "set!":
			v := e[1].(data.Symbol)
			env.Get(v).vars[v] = Eval(e[2], env)
			result = "ok"

		case "define":
			env.vars[e[1].(data.Symbol)] = Eval(e[2], env)
			result = "ok"

		case "lambda":
			result = proc{e[1], e[2], env}

		case "begin":
			for _, i := range e[1:] {
				result = Eval(i, env)
			}

		default:
			operands := e[1:]
			results := make([]data.Data, len(operands))
			for i, x := range operands {
				results[i] = Eval(x, env)
			}
			result = apply(Eval(e[0], env), results)
		}

	default:
		log.Println("EVAL - Unknown Expression Type", e)
	}
	return result
}

func apply(procedure data.Data, args []data.Data) data.Data {
	var result data.Data

	switch p := procedure.(type) {
	case func(...data.Data) data.Data:
		result = p(args...)

	case proc:
		en := &Environment{make(vars), p.en}
		switch params := p.params.(type) {
		case []data.Data:
			for i, param := range params {
				en.vars[param.(data.Symbol)] = args[i]
			}

		default:
			en.vars[params.(data.Symbol)] = args
		}
		result = Eval(p.body, en)

	default:
		log.Println("APPLY - Unknown Procedure Type", p)
	}
	return result
}

type proc struct {
	params, body data.Data
	en           *Environment
}
