package immutable

import (
	"errors"
)

// ErrBreak can be used in Range functions to stop the iteration early.
var ErrBreak = errors.New("immutable: stop iteration")
