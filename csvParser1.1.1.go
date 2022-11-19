package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	filepath := os.Args[1]
	numOfStringsToParse := os.Args[2]

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)

	pulledHeader, err := reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	terminalWidth := consoleSize()
	cellMaxLen := terminalWidth / len(pulledHeader)
	cellCapacity := cellMaxLen - 4
	drawALine(terminalWidth)

	if numOfStringsToParse == "all" {
		for {
			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			renderString(record, cellCapacity)
			//freeSpaceInEl := maxFreeSpaceInEl
			//elementTemplate := "| "
			//ouptutElement := elementTemplate

			for i := 0; i < len(record); i++ {
				pulledElement := record[i]
				for k := 0; k < len(record[i]); k++ {

					if k == len(pulledElement)-1 && freeSpaceInEl > 0 {
						ouptutElement += string(pulledElement[k])
						freeSpaceInEl--

						if freeSpaceInEl == 0 {
							ouptutElement += " |"
							fmt.Println(ouptutElement)
						} else {
							for j := 0; j < freeSpaceInEl; j++ {
								ouptutElement += " "
							}
							ouptutElement += " |"
							fmt.Println(ouptutElement)
						}
						break
					}

					if freeSpaceInEl == 0 {
						ouptutElement += " |"
						fmt.Println(ouptutElement)
						ouptutElement = ""

						for f := 0; f < i*elementMaxLen; f++ {
							ouptutElement += " "
						}

						ouptutElement += elementTemplate
						freeSpaceInEl = maxFreeSpaceInEl
					}
					ouptutElement += string(pulledElement[k])
				}

			}

		}
	}
	return
}

func consoleSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	/*heigth, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}*/

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return /*heigth,*/ width
}

func renderLine(terminalWidth int) {
	line := ""
	for i := 0; i < terminalWidth; i++ {
		line += "_"
	}
	fmt.Println(line)
	return
}

func renderString(elements []string, cellCapacity int) {
	globalLine := ""
	lineTemplate := "| "
	substringsCollection := []string{}

	for i := 0; i < len(elements); i++ {

		pushFlag := 0
		element := elements[i]
		substringsInCell := 0
		line := lineTemplate
		freeSpaceInCell := cellCapacity

		for j := 0; j < len(element); j++ {

			if freeSpaceInCell == 0 {

				if len(substringsCollection) == 0 || substringsInCell > len(substringsCollection) {
					pushFlag = 1
					globalLine += line
					substringsInCell++
					line = lineTemplate
					freeSpaceInCell := cellCapacity
					line += string(element[j])
					freeSpaceInCell--
					continue
				} else if len(substringsCollection) != 0 && substringsInCell <= len(substringsCollection) {
					pushFlag = 1
					globalLine += line
					line = substringsCollection[substringsInCell]
					freeSpaceInCell := cellCapacity
					line += string(element[j])
					freeSpaceInCell--
					continue
				}
			}

			line += string(element[j]) //!!!BASIC CASE!!!
			freeSpaceInCell--          //!!!BASIC CASE!!!
		}
		line += " |"
	}
	fmt.Println(line)
}
