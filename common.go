package immutable

import (
	"errors"
)

// ErrBreak can be used in Iter functions to stop the iteration early.
var ErrBreak = errors.New("stop iteration")
