package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

//readBlock will check for ending brackets,and if no such
// exist on the same line it will combine the line with
// the next one, and returned the combined result.
func readBlock(r *bufio.Reader) (string, error) {
	var lString string
	var l []byte
	var err error
	var tmpString string

	fmt.Println("--------------------------------------------------------------------")
	for {

		l, _, err = r.ReadLine()
		if err != nil {
			log.Println("Error: failed reading line: ", err)
			return "", err
		}

		//Checks the first 4 characters of the next line.
		peek, _ := r.Peek(4)
		peekString := strings.TrimSpace(string(peek))
		lString = strings.TrimSpace(string(l))
		//fmt.Printf("== lString ==:%v\n", lString)

		startOK := strings.HasPrefix(lString, "<")
		simpleEndOK := strings.HasSuffix(lString, ">")
		peekStartOK := strings.HasPrefix(string(peekString), "<")
		_ = fmt.Sprintln(startOK, simpleEndOK, peekStartOK)

		// In the IF below just specify when you want to run once, or break out.
		// Breaking out means read no more lines this run.
		//
		if (startOK && simpleEndOK) || //check if line contains "<" and ">", indicating complete tag line, or...
			peekStartOK { //check if next line contains "<", indicating new bracket on next line
			tmpString = fmt.Sprintf("%v %v", tmpString, lString)

			// If the finnished line don't have any brackets at all, we assume it is
			// a description, so we add new tags called <description> & </description>.
			//
			tmpString = strings.TrimSpace(tmpString)
			startOK := strings.HasPrefix(tmpString, "<")
			simpleEndOK := strings.HasSuffix(tmpString, ">")
			if !startOK && !simpleEndOK {
				tmpString = fmt.Sprintf("<description>%v</description>", tmpString)
			}

			//Since strings.Trimspace removed all carriage returns, we now add
			//a new one to the line we want to return.
			tmpString = fmt.Sprintf("%v\n", tmpString)
			break
		}
		tmpString = fmt.Sprintf("%v %v", tmpString, lString)
	}
	fmt.Println("--------------------------------------------------------------------")
	return tmpString, nil
}
