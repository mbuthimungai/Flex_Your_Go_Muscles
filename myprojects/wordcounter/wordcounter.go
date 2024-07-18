package wordcounter

import (
	"fmt"
	"os"
	"strings"
)

func WordCounter() {
	filePath := "./wordcounter/randomText.txt"
	fmt.Printf("Reading file %v\n", filePath)	
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("The error is: %v\n", err)
	}
	arrayData := strings.Fields(string(data))
	
	count := 0 
	for _, data := range(arrayData) {
		if data != " "{
			count += 1
		}		
	}
	fmt.Printf("Number of words %v\n", count)
}