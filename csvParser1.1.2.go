package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
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

	/*
		pulledHeader, err := reader.Read()

		if err != nil {
			log.Fatal(err)
		}*/

	terminalWidth := consoleSize()

	//freeSpaceInCell := cellCapacity
	//renderLine(terminalWidth)
	//theLargestElSize := calculateTheLargestElementSize(filepath)
	/*sizeOfStrArr := 0
	if theLargestElSize/cellCapacity < 1 {
		sizeOfStrArr = 1
	} else {
		sizeOfStrArr = int(math.Round(float64(theLargestElSize) / float64(cellCapacity)))
	}*/
	//renderString(pulledHeader, cellCapacity)

	if numOfStringsToParse == "all" {
		for {
			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			cellCapacity := (terminalWidth / len(record)) - 4
			numOfStringsPerRecord := maxStringsPerRecord(record, cellCapacity)
			strings := initGlobalArr(numOfStringsPerRecord)
			renderLine(terminalWidth)

			/*for _, element := range record {
				cells := initStrArr(sizeOfStrArr)
				currentSubstrCounter := 0
				for i := 0; i < len(element); i++ {
					if freeSpaceInCell == 0 {
						currentSubstrCounter++
						freeSpaceInCell = cellCapacity
					}
					cells[currentSubstrCounter] += string(element[i]) //basic action
					freeSpaceInCell--                                 //basic action
				}
				for position, item := range cells {
					placeholder := item
					cells[position] = fillStr(placeholder, cellCapacity)
				}
				for index, item := range cells {
					strings[index] += item
				}
			}*/

			for _, item := range record {
				emptyCell := renderEmptyCell(cellCapacity)
				cell := fillStrings(renderBlock(item, cellCapacity), cellCapacity)
				makeOneFromTwoV2(strings, cell, emptyCell)
			}

			//fmt.Println(strings)
			renderString(strings)
			renderLine(terminalWidth)
		}
	} else {
		counter := 0
		for {
			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}
			//counter := 0
			strNumToParse, err := strconv.Atoi(numOfStringsToParse)

			if err != nil {
				fmt.Println(`Your argument can be only an existing number of string to parse or "all" to parse the whole document`)
				break
			}

			if counter == 0 || counter == strNumToParse {
				cellCapacity := (terminalWidth / len(record)) - 4
				numOfStringsPerRecord := maxStringsPerRecord(record, cellCapacity)
				strings := initGlobalArr(numOfStringsPerRecord)
				renderLine(terminalWidth)

				for _, item := range record {
					emptyCell := renderEmptyCell(cellCapacity)
					cell := fillStrings(renderBlock(item, cellCapacity), cellCapacity)
					makeOneFromTwoV2(strings, cell, emptyCell)
				}

				renderString(strings)
				renderLine(terminalWidth)
			}

			counter++
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

func renderEmptyCell(cellCapacity int) string {
	cell := "| "
	for x := 0; x < cellCapacity; x++ {
		cell += " "
	}
	cell += " |"
	return cell
}

func renderStringOldVers(elements []string, cellCapacity int) {
	globalLine := ""
	lineTemplate := "| "
	substringsCollection := []string{}

	for i := 0; i < len(elements); i++ {
		//pushFlag := 0
		//insertFlag := 0
		element := elements[i]
		substringsInCell := 0
		line := lineTemplate
		freeSpaceInCell := cellCapacity

		for j := 0; j < len(element); j++ {

			if freeSpaceInCell == 0 {

				if len(substringsCollection) == 0 || substringsInCell > len(substringsCollection) {
					//pushFlag = 1
					line += " |"
					globalLine += line
					substringsInCell++
					substringsCollection = append(substringsCollection, line)
					line = lineTemplate
					freeSpaceInCell := cellCapacity
					line += string(element[j])
					freeSpaceInCell--
				} else if len(substringsCollection) != 0 && substringsInCell <= len(substringsCollection) {
					//insertFlag = 1
					line += " |"
					globalLine += line
					line = substringsCollection[substringsInCell]
					if len(substringsCollection) < i {
						for b := 0; b < i-1-len(substringsCollection); b++ {
							emptyCell := renderEmptyCell(cellCapacity)
							line += emptyCell
						}
					}
					line += "| "
					freeSpaceInCell := cellCapacity
					line += string(element[j])
					freeSpaceInCell--
				}
			}
			line += string(element[j])
			freeSpaceInCell--
		}

		/*if j == (len(element) - 1) {
			for g := 0; g < freeSpaceInCell; g++ {
				line += " "
			}
			line += " |"
			if pushFlag == 1 {
				substringsCollection = append(substringsCollection, line)
			} else if insertFlag == 1 {
				substringsCollection[substringsInCell] = line
			} else {
				globalLine += line
			}
		}*/
	}
	fmt.Println(globalLine)

	if len(substringsCollection) != 0 {
		for i := 0; i < len(substringsCollection); i++ {
			fmt.Println(substringsCollection[i])
		}
	}
	return
}

/*func (parsedString []string) int {

	for _, item := range parsedString {

	}
	return sizeOTheLargest
}*/

func maxStringsPerRecord(parsedString []string, cellCapacity int) int {
	sizeOTheLargestElement := 0
	for _, element := range parsedString {
		if len(element) > sizeOTheLargestElement {
			sizeOTheLargestElement = len(element)
		}
	}
	stringsPerRecord := int(math.Round(float64(sizeOTheLargestElement) / float64(cellCapacity)))
	return stringsPerRecord
}

func initGlobalArr(size int) []string {
	arr := []string{}
	for i := 0; i < size; i++ {
		arr = append(arr, "")
	}
	return arr
}

/*func fillStr(str string, cellCap int) string {
	if len(str) < cellCap {
		for i := len(str); i < cellCap; i++ {
			str += " "
		}
	}
	str += " |"
	return str
}*/

func renderBlock(element string, capacity int) []string {
	writtenData := []string{}
	currentString := ""
	spaceInCell := capacity

	for i := 0; i < len(element); i++ {

		if spaceInCell == 0 {
			writtenData = append(writtenData, currentString)
			currentString = ""
			spaceInCell = capacity

		}
		currentString += string(element[i]) //basic option
		spaceInCell--

		if i == len(element)-1 {
			writtenData = append(writtenData, currentString)
		}
	}
	return writtenData
}

func fillStrings(elements []string, capacity int) []string {
	result := []string{}

	for _, item := range elements {
		element := "| "
		element += item
		if len(item) < capacity {
			for j := len(item); j < capacity; j++ {
				element += " "
			}
		}
		element += " |"
		result = append(result, element)
	}
	return result
}

func makeOneFromTwoV2(globalArr []string, currenArr []string, emptyCell string) {
	for i := 0; i < len(globalArr); i++ {
		if i > len(currenArr)-1 {
			globalArr[i] += emptyCell
		} else {
			globalArr[i] += currenArr[i]
		}
	}
	//fmt.Println(arr1)
}

func renderString(elements []string) {
	for _, record := range elements {
		fmt.Println(record)
	}
}
