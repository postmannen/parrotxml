package main

import (
	"fmt"
	"log"
	"strings"
)

//stack will keep track of where we are working in the iteration,
type tagStack struct {
	data []string
}

func newTagStack() *tagStack {
	return &tagStack{
		//data: make([]string, 0, 100),
	}
}

//push will add another item to the end of the stack with a normal append
func (s *tagStack) push(d string) {
	s.data = append(s.data, d)
	fmt.Printf("DEBUG: PUSH on stack : %#v\n", s)
}

//pop will remove the last element of the stack
func (s *tagStack) pop() {
	//fmt.Println("DEBUG: Before pop:", s)
	last := len(s.data)
	// ---
	s.data = append(s.data[0:0], s.data[:last-1]...)

	//DEBUG BELOW
	fmt.Printf("DEBUG: POP stack:%#v\n", s)
	if len(s.data) == 0 {
		log.Println("*** STACK IS EMPTY ***")
	}

}

// =============================================================================

//findTag will check for tags at the start and end of a line
func findTag(theWord string, line string) (found bool) {
	if len(line) > 0 {
		found = strings.HasPrefix(line, theWord)
		if found {
			//fmt.Println("word found while slicing at the start of line: ", theWord)
			return true
		}
		found = strings.HasSuffix(line, theWord)
		if found {
			//fmt.Println("word found while slicing at the end of line: ", theWord)
			return true
		}
	}
	return false
}
