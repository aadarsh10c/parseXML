package parser

import (
	"fmt"
	"log"
	"os"
)

func ParseCdr(fileName string) {
	//Read the file
	data, err := os.ReadFile("CDRSample.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CDR %v", string(data))
}
