package main

import (
	"errors"
	"fmt"
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
	fmt.Printf("DEBUG: Putting on stack : %#v\n", s)
}

//pop will remove the last element of the stack
func (s *tagStack) pop() {
	//fmt.Println("DEBUG: Before pop:", s)
	last := len(s.data)
	// ---
	s.data = append(s.data[0:0], s.data[:last-1]...)
	fmt.Printf("DEBUG: After pop:%#v\n", s)

}

// =============================================================================

// =============================================================================

func printLine(line []byte) {
	//fmt.Printf("Line : %v \n Type %T\n", line, line)
	for i := 0; i < len(line); i++ {
		character := string(line[i])
		fmt.Print(character)

	}
	fmt.Println()
}

//find tag will check if there is a <project> tag in xml
func findTag(theWord string, line []byte) (found bool) {
	var tag string
	if len(line) > 0 {
		//check at the beginning of the line
		tag = string(line[0:len(theWord)])
		if tag == theWord {
			//fmt.Println("word found while slicing : ", tag)
			return true
		}

		//check at the end of the line, some tags like comments
		// end the tag on a later line with />
		end := strings.HasSuffix(string(line), theWord)
		if end {
			//fmt.Println("word found while slicing at the end of line: ", theWord)
			return true
		}
	}
	return false
}

//checkForClosingBracket
//Check for opening and closing angle bracket.
//Will return nil if both start and end bracker was found.
func checkForClosingBracket(line []byte) error {
	for i := 0; i < len(line); i++ {
		character := string(line[i])
		if character == "<" {
			ii := 0
			for {
				if string(line[ii]) == ">" {
					//fmt.Println("Found closing angle bracket at position: ", ii)
					break
				}
				if ii == len(line)-1 {

					return errors.New("Missing ending angle bracket")
				}
				ii++
			}
		}
	}
	return nil
}
