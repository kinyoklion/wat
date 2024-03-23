package wat

type Importance string

var (
	ImportanceMinor     Importance = "minor"
	ImportanceStandard  Importance = "standard"
	ImportanceImportant Importance = "important"
	ImportanceCritical  Importance = "critical"
)

type BasicDate struct {
	day   int
	month int
	year  int
}

type Entry struct {
	details    string
	date       BasicDate
	importance Importance
}
