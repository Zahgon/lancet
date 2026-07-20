package slice

func ForEachConcurrent[T any](slice []T, iteratee func(index int, item T), numThreads int) {
	_ = "STUB: not implemented"
	return
}

func MapConcurrent[T any, U any](slice []T, iteratee func(index int, item T) U, numThreads int) []U {
	_ = "STUB: not implemented"
	return nil
}

func ReduceConcurrent[T any](slice []T, initial T, reducer func(index int, item T, agg T) T, numThreads int) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func FilterConcurrent[T any](slice []T, predicate func(index int, item T) bool, numThreads int) []T {
	_ = "STUB: not implemented"
	return nil
}

func UniqueByConcurrent[T comparable](slice []T, comparator func(item T, other T) bool, numThreads int) []T {
	_ = "STUB: not implemented"
	return nil
}
