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

	terminalWidth := consoleSize()

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

			for _, item := range record {
				emptyCell := renderEmptyCell(cellCapacity)
				cell := fillStrings(renderBlock(item, cellCapacity), cellCapacity)
				makeOneFromTwoV2(strings, cell, emptyCell)
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

func maxStringsPerRecord(parsedString []string, cellCapacity int) int {
	sizeOTheLargestElement := 0
	for _, element := range parsedString {
		if len(element) > sizeOTheLargestElement {
			sizeOTheLargestElement = len(element)
		}
	}
	//stringsPerRecord := int(math.Round(float64(sizeOTheLargestElement) / float64(cellCapacity)))
	//stringsPerRecord := float64(sizeOTheLargestElement) / float64(cellCapacity)
	stringsPerRecord := float64(sizeOTheLargestElement) / float64(cellCapacity)
	if math.Remainder(stringsPerRecord, 1) > 0 {
		stringsPerRecord = math.Floor(stringsPerRecord)
		stringsPerRecord++
	}
	return int(stringsPerRecord)
}

func initGlobalArr(size int) []string {
	arr := []string{}
	for i := 0; i < size; i++ {
		arr = append(arr, "")
	}
	return arr
}

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
}

func renderString(elements []string) {
	for _, record := range elements {
		fmt.Println(record)
	}
}
