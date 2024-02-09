package goturing

import (
	"errors"
)

type TuringObject struct {
	initial_state *StateObject
}

func TuringMachine(initial_state *StateObject) (*TuringObject, error) {
	if initial_state.is_initial {
		return &TuringObject{initial_state}, nil
	} else {
		return nil, errors.New("given state is not set as initial")
	}
}