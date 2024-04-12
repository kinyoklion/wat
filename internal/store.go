package wat

import (
	"github.com/lithammer/fuzzysearch/fuzzy"

	"reflect"
	"sort"
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
		filteredEntries = Filter(filteredEntries, func(entry Entry) bool {
			return entry.Date == withOpts.date
		})
	}

	if !reflect.ValueOf(withOpts.importance).IsZero() {
		filteredEntries = Filter(filteredEntries, func(entry Entry) bool {
			return entry.Importance == withOpts.importance
		})
	}

	if !reflect.ValueOf(withOpts.text).IsZero() {
		matches := fuzzy.RankFind(withOpts.text, Select(filteredEntries, func(item Entry) string {
			return item.Details
		}))
		sort.Sort(matches)
		// TODO: Apply this ordering back to the filtered entries.
	}

	return filteredEntries
}

func (s *Store) GetEntry(id string) (value Entry, present bool) {
	for date := range s.days {
		for entry := range s.days[date].entries {
			if entry == id {
				return s.days[date].entries[entry], true
			}
		}
	}
	return Entry{}, false
}
