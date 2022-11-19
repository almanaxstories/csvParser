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
	renderLine(terminalWidth)

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
					line += " |"
					globalLine += line
					substringsInCell++
					line = lineTemplate
					freeSpaceInCell := cellCapacity
					line += string(element[j])
					freeSpaceInCell--
				} else if len(substringsCollection) != 0 && substringsInCell <= len(substringsCollection) {
					pushFlag = 1
					line += " |"
					globalLine += line
					line = substringsCollection[substringsInCell]
					freeSpaceInCell := cellCapacity
					line += string(element[j])
					freeSpaceInCell--
				}
			}

			line += string(element[j])
			freeSpaceInCell--

			if j == len(element)-1 {
				for g := 0; g < freeSpaceInCell; g++ {
					line += " "
				}
				line += " |"
				if pushFlag == 1 {
					substringsCollection[substringsInCell] = line
				} else {
					globalLine += line
					//fmt.Println(globalLine)
				}
			}
		}
	}
	fmt.Println(globalLine)
	fmt.Println(substringsCollection)
}
