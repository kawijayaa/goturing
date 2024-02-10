package goturing

import (
	"math"
	"slices"
)

type GrowingListObject struct {
	positive []rune
	negative []rune
}

func GrowingList() *GrowingListObject {
	return &GrowingListObject{make([]rune, 1), make([]rune, 0)}
}

func (gl *GrowingListObject) Set(index int, char rune) {
	if index >= 0 {
		if index < len(gl.positive) {
			gl.positive[index] = char
		} else {
			gl.positive = append(gl.positive, make([]rune, index - len(gl.positive) + 1)...)
			gl.positive[index] = char
		}
	} else {
		var adjusted_index int = int(math.Abs(float64(index))) - 1
		if adjusted_index < len(gl.negative) {
			gl.negative[adjusted_index] = char
		} else {
			gl.negative = append(gl.negative, make([]rune, adjusted_index - len(gl.negative) + 1)...)
			gl.negative[adjusted_index] = char
		}
	}
}

func (gl *GrowingListObject) Get(index int) rune {
	if index >= 0 {
		if index >= len(gl.positive) {
			return '~'
		}

		return gl.positive[index]
	} else {
		var adjusted_index int = int(math.Abs(float64(index))) - 1

		if adjusted_index >= len(gl.negative) {
			return '~'
		}

		return gl.negative[adjusted_index]
	}
}

func (gl *GrowingListObject) Slice() []rune {
	var reversed_negative []rune = make([]rune, len(gl.negative))

	copy(reversed_negative, gl.negative)
	slices.Reverse(reversed_negative)
	
	return append(reversed_negative, gl.positive...)
}

func (gl *GrowingListObject) CheckIndex(index int) bool {
	return index < len(gl.positive) && int(math.Abs(float64(index))) - 1 >= len(gl.negative)
}