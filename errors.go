package ratelimit

import (
	"fmt"
)

type QuantumError struct {
	Rate float64
}

func (e *QuantumError) Error() string {
	return fmt.Sprintf("cannot find suitable quantum for rate: %f", e.Rate)
}

type Field int

const (
	FieldFillInterval = iota
	FieldCapacity
	FieldQuantum
)

var fieldMap = map[Field]string{FieldFillInterval: "Interval", FieldCapacity:"Capacity", FieldQuantum:"Quantum"}

type ValueError struct {
	Field Field
	Value int64
}

func (e *ValueError) Error() string {
	return fmt.Sprintf("Field '%s' have incorrect value: %d", fieldMap[e.Field], e.Value)
}