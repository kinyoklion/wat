package collections

func Map[TEntry any, TOut any](items []TEntry, selector func(item TEntry) TOut) []TOut {
	outItems := make([]TOut, len(items))
	for i, item := range items {
		outItems[i] = selector(item)
	}
	return outItems
}
