package main

import (
	"errors"
	"fmt"
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
	fmt.Println("DEBUG: Putting on stack : ", s)
}

//pop will remove the last element of the stack
func (s *tagStack) pop() {
	fmt.Println("DEBUG: Before pop:", s)
	last := len(s.data)
	// ---
	s.data = append(s.data[0:0], s.data[:last-1]...)
	fmt.Println("DEBUG: After pop:", s)

}

// =============================================================================

type xmlTree struct {
}

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
		tag = string(line[0:len(theWord)])
		if tag == theWord {
			return true
		}
	}
	return false
}

//findWord looks for a word, and returns the position the last character found in slice.
// Returns zero if no word was found.
func findWord(line []byte, myWordString string) (lastPosition int) {
	//find word in []byte
	myWordByte := []byte(myWordString)
	foundWord := false

	for linePosition := 0; linePosition < len(line)-len(myWordByte); linePosition++ {
		wordPosition := 0
		for {

			//Since the iteration over the word using wordPosition as a counter will break out
			// if there is a mismatch in the matching, we can be sure that the word was found
			// if word position reaches the same value as the length of the word.
			// And we can then return the result and exit.
			if wordPosition >= len(myWordByte) {
				fmt.Println("Reached the end of the word, breaking out of word loop", linePosition, wordPosition)
				foundWord = true
				lastPosition = linePosition + wordPosition
				return lastPosition
			}

			//If there is no match break out of the loop imediatly, since there is no reason
			// to continue if one fails. Better to break out of the inner for loop and start
			// the iteration of the next charater and see if we are more lucky.
			if line[linePosition+wordPosition] != myWordByte[wordPosition] {
				break
			}

			wordPosition++
		}

		if foundWord {
			fmt.Println("Breaking out of outer loop")
			break
		}
	}
	return 0
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