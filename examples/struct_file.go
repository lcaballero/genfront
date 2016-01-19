package examples

//go:generate genfront fields --output struct_file.gen.go --template ../.files/struct_sql_tomap.fm --line $GOLINE
type Effort struct {
	Id           int
	Title        string
	Summary      string
	Description  string
	CreatedBy    int
	CreatedOn    string
	UpdatedBy    int
	UpdatedOn    string
	OwnedBy      int
	State        ScrumState
	RecordStatus RecordStatus
}

type ScrumState string
type RecordStatus int

const (
	InProgress ScrumState = "InProgress"
)
const (
	Active RecordStatus = 1
)

type User struct {
	Id       int
	Username string
	Password string
}
