package main

import (
	p "github.com/aadarsh10c/go-playground/parseXML/parser"
	// t "github.com/aadarsh10c/go-playground/parseXML/types"
)

func main() {
	// c := &t.CDR{}
	p.ParseCdr("CDRSample.xml")
}
