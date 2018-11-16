package main

import (
	"fmt"
	"log"
	"strings"
)

//chrPositions , finds the positions containing a chr in a string
//
func findChrPositions(s string, chr byte) (equalPosition []int) {
	for i := 0; i < len(s); i++ {
		//find the positions of the "=" character
		if s[i] == byte(chr) {
			//fmt.Println("DEBUG: line : ", s)
			//
			//fmt.Println("DEBUG: chr = ", string(chr), " Found chr at", equalPosition)
			equalPosition = append(equalPosition, i)
		}
	}
	return
}

//findPriorOccurance .
// Searches backwards in a string from a given positions,
// for the first occurence of a character.
//
func findPriorOccurance(s string, preChr byte, origChrPositions []int) (preChrPositions []int) {
	for _, v := range origChrPositions {
		vv := v

		//find the first space before the preceding word
		for {
			vv--

			if vv < 0 {
				log.Println("Found no space before the equal sign, reached the beginning of the line")
				break
			}
			if s[vv] == preChr {
				preChrPositions = append(preChrPositions, vv)
				break
			}
		}
	}

	//Will return the position of the prior occurance of the a the character
	return
}

//findNextOccurance .
// Searches forward in a string from a given positions,
// for the first occurence of a character after it.
// The function takes multiple positions as input,
// and will also return multiplex positions
//
func findNextOccurance(s string, preChr byte, origChrPositions []int) (nextChrPositions []int) {
	for _, v := range origChrPositions {
		vv := v

		//find the first space before the preceding word
		for {
			vv++

			if vv > len(s)-1 {
				log.Println("Found no space before the equal sign, reached the end of the line")
				break
			}

			if s[vv] == preChr {
				nextChrPositions = append(nextChrPositions, vv)
				break
			}
		}
	}

	//will return the preceding chr's positions found
	return
}

//findLettersBetween
// takes a string, and two positions given as slices as input,
// and returns a slice of string with the words found.
//
func findLettersBetween(s string, firstPositions []int, secondPositions []int) (words []string) {
	for i, v := range firstPositions {
		letters := []byte{}

		//as long as first position is lower than second position....
		for v < secondPositions[i] {
			letters = append(letters, s[v])
			v++
		}
		words = append(words, string(letters))
	}
	return
}

//getAttributes
// takes a string as input, and return the attribute names and
// values as two different slices. Reason for using slices and
// not maps are to preserve the order.
//
func (l *lexer) getAttributes() {
	//Find the positions where there is an equal sign in the string
	equalPositions := findChrPositions(l.currentLine, '=')
	preChrPositions := findPriorOccurance(l.currentLine, ' ', equalPositions)

	//==============find the word before the equal sign==============================

	//We need to add 1 to all the pre positions, since the word we're
	// looking for starts after that character.
	for i := range preChrPositions {
		preChrPositions[i]++
	}

	l.attributes.name = findLettersBetween(l.currentLine, preChrPositions, equalPositions)

	// =================find the word after the equal and between " "===========================

	nextChrPositions := findNextOccurance(l.currentLine, '"', equalPositions)
	nextNextChrPositions := findNextOccurance(l.currentLine, '"', nextChrPositions)

	//We need to add 1 to all the pre positions, since the word we're
	// looking for starts after that character.
	for i := range nextChrPositions {
		nextChrPositions[i] = nextChrPositions[i] + 1
	}

	l.attributes.value = findLettersBetween(l.currentLine, nextChrPositions, nextNextChrPositions)
}

//readCommentBlock will read another line, and add that line to the current line.
// If an end tag "/>" is found, it will break out of the for loop, and exit.
func (l *lexer) combineCommentLines() {
	newLine := l.currentLine
	for {
		newLine = fmt.Sprintf("%v %v", newLine, l.nextLine)
		if strings.HasSuffix(l.nextLine, "/>") {
			l.currentLine = newLine + "\n"
			break
		}
		// If nextLine have no close at the end, read another line pair.
		err := l.readLines()
		if err != nil {
			fmt.Println("Error: failed reading line inside readComment func: ", err)
			break
		}

	}
}
