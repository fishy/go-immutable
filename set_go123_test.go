//go:build go1.23

package immutable_test

import (
	"iter"

	"go.yhsif.com/immutable"
)

var _ iter.Seq[int] = immutable.EmptySet[int]().All()
