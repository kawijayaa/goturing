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

func (turing *TuringObject) IsDeterministic() bool {
	return is_deterministic_recursive(turing.initial_state, make([]*StateObject, 0))
}

func is_deterministic_recursive(state *StateObject, visited_states []*StateObject) bool {
	for _, visited_state := range visited_states {
		if visited_state == state {
			return true
		}
	}

	visited_states = append(visited_states, state)
	
	if len(state.transitions) == 0 {
		return true
	}

	var read_chars []rune = make([]rune, 0)

	for _, transition := range state.transitions {
		for _, char := range read_chars {
			if char == transition.read_char {
				return false
			}
		}

		read_chars = append(read_chars, transition.read_char)
	}

	is_deterministic := true
	for _, transition := range state.transitions {
		is_deterministic = is_deterministic && is_deterministic_recursive(transition.state_destination, visited_states)
		if !is_deterministic {
			return is_deterministic
		}
	}

	return is_deterministic
}

func (turing *TuringObject) Run(input string, steps bool) (Status, error) {
	if !turing.IsDeterministic() {
		return UNDEFINED, errors.New("turing machine is not deterministic")
	}
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
			return REJECTED, nil
		}
	}

	return ACCEPTED, nil
}