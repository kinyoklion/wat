package wat

import "github.com/charmbracelet/huh"

type EntryFormData struct {
	day        string
	details    string
	importance Importance
}

func EntryForm() (*huh.Form, *EntryFormData) {
	data := &EntryFormData{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("When").Options(
				huh.NewOption("Today", "today"),
				huh.NewOption("Next Business Day", "next-business"),
				huh.NewOption("Tomorrow", "tomorrow"),
			).Value(&data.day),
			huh.NewSelect[Importance]().Title("Importance").Options(
				huh.NewOption("Minor", ImportanceMinor),
				huh.NewOption("Standard", ImportanceStandard),
				huh.NewOption("Important", ImportanceImportant),
				huh.NewOption("Critical", ImportanceCritical),
			).Value(&data.importance),
			huh.NewInput().Title("Details").Value(&data.details)),
	)

	return form, data
}
