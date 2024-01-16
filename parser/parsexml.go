package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html/charset"
)

func ParseXMLtoStruct(fileName string, p interface{}) {
	//Read the file
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("error while opening file: ", err)
	}

	//converted to reader
	fileReader := bytes.NewReader(file)

	//used the reader and convert it into a decoder
	decoder := xml.NewDecoder(fileReader)

	//set the charsetReader to NewReader label function,
	//converts the content to utf-8 encoded
	decoder.CharsetReader = charset.NewReaderLabel

	//Parse the CDR xml into the struct
	err = decoder.Decode(p)
	if err != nil {
		log.Fatalf("Error while parsing: %v ", err)
	}
	fmt.Printf("Decoded %v", p)

}
