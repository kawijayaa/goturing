package goturing

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type TapeObject struct {
	pointer_index int
	tape *GrowingListObject
}

func Tape() *TapeObject {
	tape := GrowingList()
	pointer_index := 0

	return &TapeObject{pointer_index, tape}
}

func (tape *TapeObject) Set(char rune) {
	tape.tape.Set(tape.pointer_index, char)
}

func (tape *TapeObject) SetInitialState(input []rune) {
	tape.tape.positive = input
}

func (tape *TapeObject) Get() rune {
	return tape.tape.Get(tape.pointer_index)
}

func (tape *TapeObject) Move(direction Direction) error {
	switch direction.EnumIndex() {
		case LEFT: {
			tape.pointer_index--
			return nil
		}

		case STAY: {return nil}

		case RIGHT: {
			tape.pointer_index++
			return nil
		}

		default: {return errors.New("invalid direction")}
	}
}

func (tape *TapeObject) Print() {
	var tape_str []string
	var pointer_str []string

	var lower_bound int = int(math.Abs(float64(len(tape.tape.negative))))
	var upper_bound int = len(tape.tape.positive) - 1

	var print_lower_boundary int = int(math.Min(float64(tape.pointer_index), float64(lower_bound)))
	var print_upper_boundary int = int(math.Max(float64(tape.pointer_index), float64(upper_bound)))

	for i := print_lower_boundary; i <= print_upper_boundary; i++ {
		char := tape.tape.Get(i)

		if char == rune(0) {
			tape_str = append(tape_str, "~")
		} else {
			tape_str = append(tape_str, string(char))
		}

		if i == tape.pointer_index {
			pointer_str = append(pointer_str, "^")
		} else {
			pointer_str = append(pointer_str, " ")
		}
	}

	fmt.Println(strings.Join(tape_str, " "))
	fmt.Println(strings.Join(pointer_str, " "))
}