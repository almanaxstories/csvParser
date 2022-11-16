package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	filepath := os.Args[1]
	strQuantityToParse := os.Args[2]

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)

	pulledHeader, err := reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	elementMaxLen := 22
	headLine := ""

	for h := 0; h < len(pulledHeader)*elementMaxLen; h++ {
		headLine += "_"
	}
	fmt.Println(headLine)

	if strQuantityToParse == "all" {
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

	return
}
