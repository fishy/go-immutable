package immutable

import (
	"errors"
)

// ErrBreak can be used in Range functions to stop the iteration early.
var ErrBreak = errors.New("stop iteration")

// Comparable defines the key type of Map and item type of Set.
//
// It must support go's comparison operators, as defined in:
// https://golang.org/ref/spec#Comparison_operators
type Comparable interface{}
