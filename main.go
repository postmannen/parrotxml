package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type tag struct {
	name  string
	token string
}

var tagsStart = []tag{
	tag{name: "<project", token: "projectStart"},
	tag{name: "<class", token: "classStart"},
	tag{name: "<cmd", token: "cmdStart"},
	tag{name: "<comment", token: "commentStart"},
	tag{name: "<enum", token: "enumStart"},
	tag{name: "<arg", token: "argStart"},
}

var tagsEnd = []tag{
	tag{name: "</project>", token: "projectEnd"},
	tag{name: "</class>", token: "classEnd"},
	tag{name: "</cmd>", token: "cmdEnd"},
	tag{name: "/>", token: "commentEnd"},
	tag{name: "</enum>", token: "enumEnd"},
	tag{name: "</arg>", token: "argEnd"},
}

var fileName = "ardrone3.xml"

type lexer struct {
	currentLine string
	nextLine    string
	bufReader   *bufio.Reader
}

//readLines will allways read the next line, by copying to previous nextLine into currentLine,
// and then do a read from the buffer and put it into nextLine. Values are stored in the
// lexer struct.
// This means that the actual reading of the file allways will be one step ahead of currentLine
// which is the line we normally work on in the rest of the program.
func (l *lexer) readLines() error {
	ln, _, err := l.bufReader.ReadLine()
	if err != nil {
		log.Printf("Error: bufio.ReadLine: %v\n", err)
	}
	l.currentLine = l.nextLine
	l.nextLine = strings.TrimSpace(string(ln))
	return err
}

func main() {
	//Open file for reading
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("Error: os.Open: %v\n", err)
	}
	defer f.Close()

	lex := &lexer{
		bufReader: bufio.NewReader(f),
	}

	lineNR := 1

	// =================Iterate and find=====================

	//Iterate the file and the xml data, and parse values. Create a stack to use.
	// The nice thing about using a stack for your tags found is that you will know
	// if there was a closing tag for each start tag.
	tagStack := newTagStack()
	doneReading := false
	for {
		err := lex.readLines()
		if err != nil {
			log.Println("Error: Reading lines from buffer: ", err)
			doneReading = true
		}

		var foundTag bool

		//Look for start tag.
		for i := range tagsStart {
			foundTag = findTag(tagsStart[i].name, lex.currentLine)
			if foundTag {
				attributeNames, attributeValues := getAttributes(string(lex.currentLine))
				fmt.Println("-----------------------------------------------------------------")
				tagStack.push(tagsStart[i].token)
				fmt.Println("--- Tag: ", tagsStart[i].token)
				fmt.Printf("--- Attributes: name : %v, value %v \n", attributeNames, attributeValues)
			}
		}

		//Look for end tag.
		for i := range tagsEnd {
			foundTag = findTag(tagsEnd[i].name, lex.currentLine)
			if foundTag {
				tagStack.pop()
				fmt.Println(tagsEnd[i].token)
			}
		}

		lineNR++
		if doneReading {
			break
		}
	}

}
