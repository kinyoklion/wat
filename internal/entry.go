package wat

type Importance string

var (
	ImportanceMinor     Importance = "minor"
	ImportanceStandard  Importance = "standard"
	ImportanceImportant Importance = "important"
	ImportanceCritical  Importance = "critical"
)

type State string

var (
	StateTodo State = "todo"
	StateDone State = "done"
)

type BasicDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Entry struct {
	ID         string     `json:"id"`
	Details    string     `json:"details"`
	Date       BasicDate  `json:"date"`
	Importance Importance `json:"importance"`
	State      State      `json:"state"`
}
