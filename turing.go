package goturing

import (
	"errors"
	"fmt"
)

type TuringObject struct {
	initial_state *StateObject
	tape *TapeObject
}

func TuringMachine(initial_state *StateObject) (*TuringObject, error) {
	if initial_state.is_initial {
		return &TuringObject{initial_state, Tape()}, nil
	} else {
		return nil, errors.New("given state is not set as initial")
	}
}

func (turing *TuringObject) Run(input string, steps bool) Status {
	turing.tape.SetInitialState([]rune(input))

	current_state := turing.initial_state

	if steps {
		fmt.Println(current_state.name)
		turing.tape.Print()
	}

	for !current_state.is_final {
		found := false
		for _, transition := range current_state.transitions {
			pointed_rune := turing.tape.Get()

			if transition.read_char == pointed_rune {
				found = true
				turing.tape.Set(transition.write_char)
				turing.tape.Move(transition.direction)

				current_state = transition.state_destination
				break
			}
		}

		if found {
			if steps {
				fmt.Println(current_state.name)
				turing.tape.Print()
			}
		} else {
			return REJECTED
		}
	}

	return ACCEPTED
}