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

	//=================INITALIZATION===================

	//Open file for reading
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
		tag{name: "<enum", token: "enumStart"},
	}

	tagsEnd := []tag{
		tag{name: "</project>", token: "projectEnd"},
		tag{name: "</class>", token: "classEnd"},
		tag{name: "</cmd>", token: "cmdEnd"},
		tag{name: "/>", token: "commentEnd"},
		tag{name: "</enum>", token: "enumEnd"},
	}

	// =================Iterate and find=====================

	//Iterate the file and the xml data, and parse values.
	// create a stack to use.
	// The nice thing about using a stack for your tags found
	// is that you will know if there was a closing tag for
	// each start tag.
	tagStack := newTagStack()
	for {
		//read a line
		readLine, _, err := fReader.ReadLine()
		if err != nil {
			log.Printf("Error: bufio.ReadLine: %v\n", err)
			break
		}

		//Remove leading spaces in the current line
		line := strings.TrimSpace(string(readLine))
		//printLine(line)

		// -----------------------Do the actual iteration-------------------

		foundTag := false

		//Look for start tag.
		for i := range tagsStart {
			foundTag = findTag(tagsStart[i].name, line)
			if foundTag {
				attributeNames, attributeValues := getAttributes(string(line))
				fmt.Printf("--- Attributes : name : %v, value %v \n", attributeNames, attributeValues)

				tagStack.push(tagsStart[i].token)
				fmt.Println(tagsStart[i].token)
			}
		}

		//Look for end tag.
		for i := range tagsEnd {
			foundTag = findTag(tagsEnd[i].name, line)
			if foundTag {
				tagStack.pop()
				fmt.Println(tagsEnd[i].token)
			}
		}

		lineNR++
	}

}
