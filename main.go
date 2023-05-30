package main

func Chunk[T any](elms []T, size int) [][]T {
	chunks := make([][]T, 0, (len(elms)+size-1)/size)
	for size < len(elms) {
		elms, chunks = elms[size:], append(chunks, elms[0:size:size])
	}
	return append(chunks, elms)
}
