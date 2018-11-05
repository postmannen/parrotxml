package main

import (
	"fmt"
	"log"
)

//findWord looks for a word, and returns the position the last character found in slice.
// Returns zero if no word was found.
// This function can be used for finding tag attributes like "some=", if found it will
// return the position if the equals sign "=" in the slice, which can be handy since you
// know the attribute is to the left of the equal, and the value is to the right.
//
// THIS FUNCTION IS PROBABLY REPLACED BY OTHERS, AND NOT NEEDED ANYMORE.
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

//chrPositions , finds the positions containing a chr in a string
//
func findChrPositions(s string, chr byte) (equalPosition []int) {
	for i := 0; i < len(s); i++ {
		//find the positions of the "=" character
		if s[i] == byte(chr) {
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
func getAttributes(s string) (attributeNames []string, attributeValues []string) {
	//Find the positions where there is an equal sign in the string
	equalPositions := findChrPositions(s, '=')
	preChrPositions := findPriorOccurance(s, ' ', equalPositions)

	//==============find the word before the equal sign==============================

	//We need to add 1 to all the pre positions, since the word we're
	// looking for starts after that character.
	for i := range preChrPositions {
		preChrPositions[i]++
	}

	attributeNames = findLettersBetween(s, preChrPositions, equalPositions)

	// =================find the word after the equal and between " "===========================

	nextChrPositions := findNextOccurance(s, '"', equalPositions)
	nextNextChrPositions := findNextOccurance(s, '"', nextChrPositions)

	//We need to add 2 to all the pre positions, since the word we're
	// looking for starts after that character.
	for i := range nextChrPositions {
		nextChrPositions[i] = nextChrPositions[i] + 1
	}

	attributeValues = findLettersBetween(s, nextChrPositions, nextNextChrPositions)
	return
}
