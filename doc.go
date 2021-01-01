// Package immutable provides immutable data structures (map/set/list).
//
// Note that immutable map/set/list only guarantee the immutability of the
// container it self, not the content inside.
// For example if you are using a immutable list of pointers,
// you are guaranteed that you always get the same pointer with the same index,
// but the content pointer points to might be changed by others shared the same
// immutable list.
package immutable // import "go.yhsif.com/immutable"
