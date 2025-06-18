package slices

func Map[T any, E any](slice []T, mapFunc func(T) E) []E {
	var result []E
	for _, item := range slice {
		result = append(result, mapFunc(item))
	}
	return result
}

func Filter[T any](slice []T, filterFunc func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if filterFunc(item) {
			result = append(result, item)
		}
	}
	return result
}
