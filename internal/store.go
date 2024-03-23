package wat

type entryOptions struct {
	date       BasicDate
	importance Importance
	text       string
}

type EntryOption func(options *entryOptions)

func WithDate(date BasicDate) {

}

func WithImportance(importance Importance) {

}

func WithText(text string) {

}

type dateEntry struct {
	entries []Entry
	dirty   bool
}

type Store struct {
	entries map[BasicDate][]Entry
}

func (s *Store) AddEntry(entry Entry) {

}

func (s *Store) GetEntry(options ...EntryOption) {
	withOpts := entryOptions{}
	for _, opt := range options {
		opt(&withOpts)
	}
}
