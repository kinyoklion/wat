package wat

import (
	"slices"
	"wat/algorithms"
	"wat/collections"

	"reflect"
)

type entryOptions struct {
	date       BasicDate
	importance Importance
	text       string
	state      State
}

type EntryOption func(options *entryOptions)

func WithDate(date BasicDate) EntryOption {
	return func(options *entryOptions) {
		options.date = date
	}
}

func WithImportance(importance Importance) EntryOption {
	return func(options *entryOptions) {
		options.importance = importance
	}
}

func WithText(text string) EntryOption {
	return func(options *entryOptions) {
		options.text = text
	}
}

func WithState(state State) EntryOption {
	return func(options *entryOptions) {
		options.state = state
	}
}

type Store struct {
	entries []Entry
}

func (s *Store) AddEntry(entry Entry) {
	s.entries = append(s.entries, entry)
}

func (s *Store) GetEntries(options ...EntryOption) []Entry {
	withOpts := entryOptions{}
	for _, opt := range options {
		opt(&withOpts)
	}

	var filteredEntries []Entry
	for _, entry := range s.entries {
		filteredEntries = append(filteredEntries, entry)
	}

	if !reflect.ValueOf(withOpts.date).IsZero() {
		filteredEntries = collections.Filter(filteredEntries, func(entry Entry) bool {
			return entry.Date == withOpts.date
		})
	}

	if !reflect.ValueOf(withOpts.importance).IsZero() {
		filteredEntries = collections.Filter(filteredEntries, func(entry Entry) bool {
			return entry.Importance == withOpts.importance
		})
	}

	if !reflect.ValueOf(withOpts.state).IsZero() {
		filteredEntries = collections.Filter(filteredEntries, func(entry Entry) bool {
			return entry.State == withOpts.state
		})
	}

	if !reflect.ValueOf(withOpts.text).IsZero() {
		slices.SortFunc(filteredEntries, func(a, b Entry) int {
			return algorithms.LevenshteinCmp(withOpts.text, a.Details, b.Details)
		})
	}

	return filteredEntries
}
