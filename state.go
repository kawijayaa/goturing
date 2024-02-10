package goturing

import (
	"errors"
)

type StateObject struct {
	name string
	transitions []*TransitionObject
	is_final bool
	is_initial bool
}

func State(name string, is_final bool, is_inital bool) *StateObject {
	return &StateObject{name, make([]*TransitionObject, 0), is_final, is_inital}
}

func (state *StateObject) AddTransition(transition *TransitionObject) error {
	if !state.is_final {
		state.transitions = append(state.transitions, transition)
		return nil
	}

	return errors.New("final state cannot have any transition")
}