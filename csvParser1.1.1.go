package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	filepath := os.Args[1]
	//strQuantityToParse := os.Args[2]

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)

	pulledHeader, err := reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	width := consoleSize()
	elementMaxLen := width / len(pulledHeader)
	drawALine(width)
	//fmt.Println("\n")
	i := 0
	for i < 6 {

		renderBlockEmptyLine(len(pulledHeader), elementMaxLen-4)

		drawALine(width)
		fmt.Println("\n")
		i++
	}
	drawALine(width)

	/*headLine := ""

	for h := 0; h < len(pulledHeader)*elementMaxLen; h++ {
		headLine += "_"
	}
	fmt.Println(headLine)*/

	/*if strQuantityToParse == "all" {
		for {
			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			maxFreeSpaceInEl := 18
			freeSpaceInEl := maxFreeSpaceInEl
			elementTemplate := "| "
			ouptutElement := elementTemplate

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

	bottomLine := ""

	for x := 0; x < len(pulledHeader)*elementMaxLen; x++ {
		bottomLine += "_"
	}
	fmt.Println(bottomLine)
	width := consoleSize()
	drawALine(width)*/

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

func drawALine(terminalWidth int) {
	line := ""
	for i := 0; i < terminalWidth; i++ {
		line += "_"
	}
	fmt.Println(line)
	return
}

func renderBlockEmptyLine(numOfCells int, symbolsInCell int) {
	line := ""
	for i := 0; i < numOfCells; i++ {
		line += "| "
		for j := 0; j < symbolsInCell; j++ {
			line += " "
		}
		line += " |"
	}
	fmt.Println(line)
}
