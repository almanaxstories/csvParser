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

	OriginalTerminalWidth := consoleSize()

	if numOfStringsToParse == "all" {
		for {
			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			cellCapacity := (OriginalTerminalWidth / len(record)) - 4
			terminalWidth := calcDynamicTerminalWidth(cellCapacity+4, len(record))
			numOfStringsPerRecord := maxStringsPerRecord(record, cellCapacity)
			strings := makeASlice(numOfStringsPerRecord)
			renderLine(terminalWidth)

			for _, item := range record {
				emptyCell := renderEmptyCell(cellCapacity)
				block := makeABlock(sliceAString(item, cellCapacity), cellCapacity)
				mergeArrays(strings, block, emptyCell)
			}
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

			strNumToParse, err := strconv.Atoi(numOfStringsToParse)

			if err != nil {
				fmt.Println(`Your argument can be only an existing number of string to parse or "all" to parse the whole document`)
				break
			}

			if counter == 0 || counter == strNumToParse {
				cellCapacity := (OriginalTerminalWidth / len(record)) - 4
				terminalWidth := calcDynamicTerminalWidth(cellCapacity+4, len(record))
				numOfStringsPerRecord := maxStringsPerRecord(record, cellCapacity)
				strings := makeASlice(numOfStringsPerRecord)
				renderLine(terminalWidth)

				for _, item := range record {
					emptyCell := renderEmptyCell(cellCapacity)
					block := makeABlock(sliceAString(item, cellCapacity), cellCapacity)
					mergeArrays(strings, block, emptyCell)
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
		line += "-"
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

func maxStringsPerRecord(parsedString []string, cellCapacity int) int {
	sizeOTheLargestElement := 0
	for _, element := range parsedString {
		if len(element) > sizeOTheLargestElement {
			sizeOTheLargestElement = len(element)
		}
	}
	stringsPerRecord := math.Ceil(float64(sizeOTheLargestElement) / float64(cellCapacity))
	return int(stringsPerRecord)
}

func makeASlice(size int) []string {
	arr := []string{}
	for i := 0; i < size; i++ {
		arr = append(arr, "")
	}
	return arr
}

/*func renderBlock(element string, capacity int) []string {
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
	}
	writtenData = append(writtenData, currentString)
	return writtenData
}*/

func makeABlock(elements []string, capacity int) []string {
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

func mergeArrays(globalArr []string, currenArr []string, emptyCell string) []string {
	for i := 0; i < len(globalArr); i++ {
		if i > len(currenArr)-1 {
			globalArr[i] += emptyCell
		} else {
			globalArr[i] += currenArr[i]
		}
	}
	return globalArr
}

func renderString(elements []string) {
	for _, record := range elements {
		fmt.Println(record)
	}
}

func calculateCellCapacity(terminalWidth int, recordSize int) {
	//template
}

func sliceAString(aString string, capacity int) []string {
	numOfSubstrings := int(math.Ceil((float64(len(aString)) / float64(capacity))))
	block := []string{}
	startIndex := 0
	for i := 0; i < numOfSubstrings; i++ {
		if i == numOfSubstrings-1 {
			block = append(block, aString[startIndex:len(aString)])
			break
		}
		block = append(block, aString[startIndex:startIndex+capacity])
		startIndex += capacity
	}

	return block
}

func calcDynamicTerminalWidth(cellCapacity int, lenOfRec int) int {
	return cellCapacity * lenOfRec
}
