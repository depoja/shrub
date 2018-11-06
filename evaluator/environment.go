package evaluator

import (
	"shrub/data"
	"shrub/parser"
)

// A map of variables, custom defined or primitive procedures
type vars map[data.Symbol]data.Data

// An Environment containing the variables and a reference to the outer Environment
type Environment struct {
	vars
	outer *Environment
}

// Get a variable from the current environment or the outer ones
func (e *Environment) Get(s data.Symbol) *Environment {
	if _, ok := e.vars[s]; ok {
		return e
	}
	return e.outer.Get(s)
}

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	var env Environment
	env = Environment{
		vars{
			"+":      add,
			"-":      sub,
			"*":      mult,
			"/":      div,
			"<=":     lteq,
			"equal?": eq,
			"cons":   construct,
			"car":    head,
			"cdr":    tail,
			"list":   Eval(parser.Parse("(lambda z z)"), &env),
		},
		nil}
	return &env
}
