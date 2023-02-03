package log

import (
	"errors"
	"fmt"
)

// ErrWrongFormat is thrown when the source file input is in an incorrect format
var ErrWrongFormat = errors.New("wrong source format")

// ErrNoPath is thrown when there is no path from src to dest
var ErrNoPath = errors.New("no path found")

// ErrMixMapping is thrown when there is a mixture of integers and strings in the input file
var ErrMixMapping = errors.New("Potential mixing of integer and string vertex ID's :" + ErrWrongFormat.Error())

// ErrLoopDetected is thrown when a loop is detected, causing the distance to go
// to inf (or -inf), or just generally loop forever
var ErrLoopDetected = errors.New("infinite loop detected")

// NewErrLoop generates a new error with details for loop error
func NewErrLoop(a, b int) error {
	return errors.New(fmt.Sprint(ErrLoopDetected.Error(), "From vertex '", a, "' to vertex '", b, "'"))
}
