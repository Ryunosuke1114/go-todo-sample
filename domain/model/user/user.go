package user

type User struct {
	Status AccessStatus
}
type AccessStatus int

const (
	fullAccess = iota
	readOnly   = iota
	canEdit    = iota
	noAccess   = iota
)
