package goturing

type StateObject struct {
	name string
	transitions []TransitionStruct
	is_final bool
	is_initial bool
}

func State(name string, is_final bool, is_inital bool) *StateObject {
	return &StateObject{name, make([]TransitionStruct, 0), is_final, is_inital}
}