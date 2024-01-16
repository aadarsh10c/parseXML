package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aadarsh10c/go-playground/parseXML/types"
)

func TestParseXML(t *testing.T) {
	p := &types.CDR{}

	ParseXMLtoStruct("../CDRSample.xml", p)
	t.Run("parse xml get the parent tag name", func(t *testing.T) {
		want := "CDR"

		compareString(p.XMLName.Local, want, t)
	})
	t.Run("parse CRD attributes from  xml", func(t *testing.T) {
		want := []string{
			"CDR",
			"test",
			"ga1-tas-0.ga1-tas.hpetas-lab-bel01-csde-ns01.svc.cluster.local",
			"2",
			"ScifSipService",
			"1",
			"scifsession_SIP_SERVICE_examplehost_1_examplehost_1_1_1",
			"2023-08-08 12:51:47.665",
			"2023-08-08 12:51:56.291",
		}
		fmt.Printf("before calling...")
		VerifyStruct(p, want, t)

	})
	t.Run("parse xml extract msec in call struct", func(t *testing.T) {
		want := 1
		got := p.Call.Msec
		if got != want {
			t.Errorf("Got %d , wanted %d", got, want)
		}
	})
	t.Run("parse xml networkCallType attribute in call struct", func(t *testing.T) {
		want := "ORIGINATING"
		got := p.Call.NetworkCallType
		compareString(got, want, t)
	})

	t.Run("parse xml networkCallType attribute in call struct", func(t *testing.T) {
		want := "ORIGINATING"
		got := p.Call.NetworkCallType
		compareString(got, want, t)
	})
	t.Run("Parse appLogInviteFwd msec ", func(t *testing.T) {
		want := []int{74, 74}
		appLogLen := len(p.Call.AppLogInviteFwd)
		got := make([]int, appLogLen)
		for i := 0; i < appLogLen; i++ {
			got[i] = p.Call.AppLogInviteFwd[i].Msec
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
	t.Run("Parse appLogInviteFwd conract ", func(t *testing.T) {
		want := []string{"BTS_1", "BTS_2"}
		appLogLen := len(p.Call.AppLogInviteFwd)
		got := make([]string, appLogLen)
		for i := 0; i < appLogLen; i++ {
			got[i] = p.Call.AppLogInviteFwd[i].Contract
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})

}

func VerifyStruct(x interface{}, want []string, t *testing.T) {
	val := reflect.ValueOf(x)

	//extract the underlying
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	//loop through want string and compae it with the struct
	for i := 1; i < len(want); i++ {
		compareString(val.Field(i).String(), want[i], t)
	}

}

func compareString(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %q , wanted %q", got, want)
	}
}
