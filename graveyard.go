package main

import (
	"unsafe"
)

// GetChunkSize() -> returns the amount of elements for chunk or 1
// if size of T > 521, (unused)
// c++ 512 bytes implementation
func GetChunkSize[T any](val T) uintptr {
	return max(512/unsafe.Sizeof(*new(T)), 1)
}
