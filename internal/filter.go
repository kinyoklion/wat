package wat

func Filter[T any](slice []T, predicate func(T) bool) (ret []T) {
	for _, item := range slice {
		if predicate(item) {
			ret = append(ret, item)
		}
	}
	return
}
