// The declarations of builtin (primitive) procedures

package evaluator

import (
	"reflect"
	"shrub/data"
)

// TODO: Remove type casting from these primitives

// Performs addition on a list
func add(d ...data.Data) data.Data {
	sum := d[0].(data.Number) // Set the sum to the first operand
	for _, num := range d[1:] {
		sum += num.(data.Number)
	}
	return sum
}

// Performs substration on a list
func sub(d ...data.Data) data.Data {
	diff := d[0].(data.Number) // Set the diff to the first operand
	for _, num := range d[1:] {
		diff -= num.(data.Number)
	}
	return diff
}

// Performs multiplication on a list
func mult(d ...data.Data) data.Data {
	prod := d[0].(data.Number) // Set the prod to the first operand
	for _, num := range d[1:] {
		prod *= num.(data.Number)
	}
	return prod
}

// Performs division on a list
func div(d ...data.Data) data.Data {
	result := d[0].(data.Number) // Set the result to the first operand
	for _, num := range d[1:] {
		result /= num.(data.Number)
	}
	return result
}

// TODO: Add support for non-boolean operands
func and(d ...data.Data) data.Data {
	result := d[0].(data.Boolean) // Set the result to the first operand
	for _, op := range d[1:] {
		result = result && op.(data.Boolean)
	}
	return result
}

// TODO: Add support for non-boolean operands
func or(d ...data.Data) data.Data {
	result := d[0].(data.Boolean) // Set the result to the first operand
	for _, op := range d[1:] {
		result = result || op.(data.Boolean)
	}
	return result
}

// TODO: Add support for non-boolean operands
func not(d ...data.Data) data.Data {
	value := d[0].(data.Boolean)
	return !value
}

// Compares (less than equal)
func lteq(d ...data.Data) data.Data {
	return d[0].(data.Number) <= d[1].(data.Number)
}

// Compares equality (deep)
func eq(d ...data.Data) data.Data {
	return reflect.DeepEqual(d[0], d[1])
}

// Returns the first element of a list
func head(d ...data.Data) data.Data {
	return d[0].([]data.Data)[0]
}

// Returns the tail of a list
func tail(d ...data.Data) data.Data {
	return d[0].([]data.Data)[1:]
}

// Creates a list from two arguments
func construct(d ...data.Data) data.Data {
	switch car := d[0]; cdr := d[1].(type) {
	// If first argument is a slice, append the second
	case []data.Data:
		return append([]data.Data{car}, cdr...)

	// Otherwise return a slice with both
	default:
		return []data.Data{car, cdr}
	}
}
