# goturing (WIP)

A turing machine implementation built with Go.

Signatures

```go
TuringMachine(initial_state *StateObject) (*TuringObject, error)
TuringObject.Run(input string, steps bool) Status

State(name string, is_final bool, is_inital bool) *StateObject
StateObject.AddTransition(transition *TransitionObject) error

Transition(read_char rune, write_char rune, direction Direction, state_destination *StateObject) *TransitionObject

Direction(LEFT = -1, STAY = 0, RIGHT = 1)

Status(REJECTED = -1, UNDEFINED = 0, ACCEPTED = 1)
```

Example usage:
```go
package main

import (
	"github.com/kawijayaa/goturing"
	"fmt"
)

func main() {
	q0 := goturing.State("q0", false, true)
	q1 := goturing.State("q1", false, false)
	q2 := goturing.State("q2", false, false)
	q3 := goturing.State("q3", false, false)
	q4 := goturing.State("q4", false, false)
	q5 := goturing.State("q5", true, false)
	q6 := goturing.State("q6", false, false)

	q0.AddTransition(goturing.Transition('+', '+', goturing.RIGHT, q0))
	q0.AddTransition(goturing.Transition('1', 'A', goturing.RIGHT, q1))

	q1.AddTransition(goturing.Transition('1', '1', goturing.RIGHT, q1))
	q1.AddTransition(goturing.Transition('+', '+', goturing.RIGHT, q1))
	q1.AddTransition(goturing.Transition('=', '=', goturing.RIGHT, q1))
	q1.AddTransition(goturing.Transition('A', 'A', goturing.LEFT, q2))
	q1.AddTransition(goturing.Transition('~', '~', goturing.LEFT, q2))

	q2.AddTransition(goturing.Transition('+', '+', goturing.LEFT, q2))
	q2.AddTransition(goturing.Transition('1', 'A', goturing.LEFT, q6))

	q3.AddTransition(goturing.Transition('1', '1', goturing.LEFT, q3))
	q3.AddTransition(goturing.Transition('+', '+', goturing.LEFT, q3))
	q3.AddTransition(goturing.Transition('=', '=', goturing.LEFT, q3))
	q3.AddTransition(goturing.Transition('A', 'A', goturing.RIGHT, q0))
	q3.AddTransition(goturing.Transition('~', '~', goturing.RIGHT, q0))

	q4.AddTransition(goturing.Transition('A', 'A', goturing.STAY, q5))

	q6.AddTransition(goturing.Transition('1', '1', goturing.LEFT, q3))
	q6.AddTransition(goturing.Transition('=', '=', goturing.LEFT, q4))

	tm, _ := goturing.TuringMachine(q0)
	result := tm.Run("11+1=111", true)

	fmt.Println(result)
}
```

Example output (with steps enabled):
```
q0
1 1 + 1 = 1 1 1
^              
q1
A 1 + 1 = 1 1 1
  ^            
q1
A 1 + 1 = 1 1 1
    ^          
q1
A 1 + 1 = 1 1 1
      ^        
q1
A 1 + 1 = 1 1 1
        ^      
q1
A 1 + 1 = 1 1 1
          ^    
q1
A 1 + 1 = 1 1 1
            ^  
q1
A 1 + 1 = 1 1 1
              ^
q1
A 1 + 1 = 1 1 1 ~
                ^
q2
A 1 + 1 = 1 1 1 ~
              ^  
q6
A 1 + 1 = 1 1 A ~
            ^    
q3
A 1 + 1 = 1 1 A ~
          ^      
q3
A 1 + 1 = 1 1 A ~
        ^        
q3
A 1 + 1 = 1 1 A ~
      ^          
q3
A 1 + 1 = 1 1 A ~
    ^            
q3
A 1 + 1 = 1 1 A ~
  ^              
q3
A 1 + 1 = 1 1 A ~
^                
q0
A 1 + 1 = 1 1 A ~
  ^              
q1
A A + 1 = 1 1 A ~
    ^            
q1
A A + 1 = 1 1 A ~
      ^          
q1
A A + 1 = 1 1 A ~
        ^        
q1
A A + 1 = 1 1 A ~
          ^      
q1
A A + 1 = 1 1 A ~
            ^    
q1
A A + 1 = 1 1 A ~
              ^  
q2
A A + 1 = 1 1 A ~
            ^    
q6
A A + 1 = 1 A A ~
          ^      
q3
A A + 1 = 1 A A ~
        ^        
q3
A A + 1 = 1 A A ~
      ^          
q3
A A + 1 = 1 A A ~
    ^            
q3
A A + 1 = 1 A A ~
  ^              
q0
A A + 1 = 1 A A ~
    ^            
q0
A A + 1 = 1 A A ~
      ^          
q1
A A + A = 1 A A ~
        ^        
q1
A A + A = 1 A A ~
          ^      
q1
A A + A = 1 A A ~
            ^    
q2
A A + A = 1 A A ~
          ^      
q6
A A + A = A A A ~
        ^        
q4
A A + A = A A A ~
      ^          
q5
A A + A = A A A ~
      ^          
ACCEPTED
```
