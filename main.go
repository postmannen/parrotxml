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

func main() {

	//------------------------------

	fileName := "ardrone3.xml"
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("Error: os.Open: %v\n", err)
	}
	defer f.Close()

	//bufio lets us read files line by line
	fReader := bufio.NewReader(f)

	//Start with the first line
	lineNR := 1

	tagsStart := []tag{
		tag{name: "<project", token: "projectStart"},
		tag{name: "<class", token: "classStart"},
		tag{name: "<cmd", token: "cmdStart"},
		tag{name: "<comment", token: "commentStart"},
	}

	tagsEnd := []tag{
		tag{name: "</project>", token: "projectEnd"},
		tag{name: "</class>", token: "classEnd"},
		tag{name: "</cmd>", token: "cmdEnd"},
		tag{name: "/>", token: "commentEnd"},
	}

	//Iterate the file and the xml data, and parse values.
	//create a stack to use
	tagStack := newTagStack()
	for {
		//read a line
		line, _, err := fReader.ReadLine()
		if err != nil {
			log.Printf("Error: bufio.ReadLine: %v\n", err)
			break
		}

		//Remove leading spaces in the current line
		tmpLine := strings.TrimSpace(string(line))
		line = []byte(tmpLine)
		//printLine(line)

		//Look for all the start tags, and if found, add to stack
		for i := range tagsStart {
			//Look for the start tag called <project>
			found := findTag(tagsStart[i].name, line)
			if found {
				tagStack.push(tagsStart[i].token)
				fmt.Println(tagsStart[i].token)
			}
		}

		//Look for all the end tags, and if found, add to stack
		for i := range tagsEnd {
			//Look for the start tag called <project>
			found := findTag(tagsEnd[i].name, line)
			if found {
				tagStack.pop()
				fmt.Println(tagsEnd[i].token)
			}
		}

		lineNR++
	}

}
