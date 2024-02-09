package goturing

type Direction int

const (
	LEFT = iota - 1
	STAY
	RIGHT
)

func (direction Direction) String() string {
	return [...]string{"LEFT", "STAY", "RIGHT"}[direction + 1]
}

func (direction Direction) EnumIndex() int {
	return int(direction)
}