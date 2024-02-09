package goturing

type Status int

const (
	REJECTED = iota - 1
	UNDEFINED
	ACCEPTED 
)

func (status Status) String() string {
	return [...]string{"REJECTED", "UNDEFINED", "ACCEPTED"}[status + 1]
}

func (status Status) EnumIndex() int {
	return int(status)
}