package data

// Data is the interface implemented by all the data types in shrub (symbols, numbers, expression, procedures, lists etc.)
type Data interface{}

// Symbol is represented as a string
type Symbol string

// Number is represented as a float64
type Number float64
