package goturing

type TransitionStruct struct {
	read_char rune
	write_char rune
	direction Direction
	state_destination *StateObject
}

func Transition(read_char rune, write_char rune, direction Direction, state_destination *StateObject) *TransitionStruct {
	return &TransitionStruct{read_char, write_char, direction, state_destination}
}