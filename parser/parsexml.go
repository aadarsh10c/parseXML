package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/aadarsh10c/go-playground/parseXML/types"
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


func ParseXMLtoStruct2( xmlEncodedData string , p *types.CDR){

	// URL-decode the data
	decodedData, err := url.QueryUnescape(xmlEncodedData)
	if err != nil {
		fmt.Printf("Error decoding URL: %v\n", err)
		return
	}
	fmt.Println("Decoded data: ",decodedData)
	err = xml.Unmarshal([]byte(decodedData), p)
	if err != nil {
		fmt.Printf("Error unmarshaling XML: %v\n", err)
		return
	}

	fmt.Printf("Parsed XML: %+v\n", (*p))
}
