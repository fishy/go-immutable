package immutable

import (
	"errors"
)

// ErrBreak can be used in Iter functions to stop the iteration early.
var ErrBreak = errors.New("stop iteration")

// Hashable defines the key type of Map and item type of Set.
//
// It must be hashable.
type Hashable interface{}
