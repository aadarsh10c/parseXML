package parser

import (
	"encoding/xml"
	"testing"
)

func TestParseXML(t *testing.T) {
	data := `
	<person>
		<name>Adarsh</name>
	</person>
	`
	data2 := `
	<appLogRoute msec="518">
         <dial_code msec="523">*</dial_code>
    </appLogRoute>
	`
	t.Run("parse name from  xml", func(t *testing.T) {
		p := &Person{}
		err := xml.Unmarshal([]byte(data), p)
		if err != nil {
			t.Fatalf("Error while parsing: %v ", err)
		}
		got := p.Name
		want := "Adarsh"
		compareString(got, want, t)
	})
	t.Run("Parse the given xml file ", func(t *testing.T) {

	})
}

func compareString(got string, want string, t *testing.T) {
	if got != want {
		t.Fatalf("Got %q , wanted %q", got, want)
	}
}
