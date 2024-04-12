package wat

func Values[Tk comparable, Tv any](input map[Tk]Tv) []Tv {
	list := make([]Tv, len(input))
	index := 0
	for _, value := range input {
		list[index] = value
		index++
	}
	return list
}
