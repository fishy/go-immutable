package immutable

import (
	"errors"
)

// ErrBreak can be used in Range functions to stop the iteration early.
var ErrBreak = errors.New("immutable: stop iteration")

// DropOK is a helper function for Map.Get when you want to drop/ignore the
// ok return.
func DropOK[T any](value T, ok bool) T {
	return value
}
