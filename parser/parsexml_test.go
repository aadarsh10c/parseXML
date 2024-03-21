package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aadarsh10c/go-playground/parseXML/types"
)

func TestParseXML(t *testing.T) {
	p := &types.CDR{}
	xmlEncodedData := `%3C%3Fxml+version%3D%221.0%22+encoding%3D%22ISO-8859-1%22+standalone%3D%22no%22%3F%3E%0A%3CCDR+xmlns%3Adebug%3D%22http%3A%2F%2Fwww.hp.com%2Focmp%2F2004%2F03%2Fcdr-debug%22+xmlns%3Alog%3D%22http%3A%2F%2Fwww.hp.com%2Focmp%2F2004%2F03%2Fcdr-log%22+contextversion%3D%221%22+host%3D%22crf-tas-dev-b2bua-tas-sipnl-0%22+level%3D%223%22+scifservicename%3D%22CrfIldTasService%22+seeinstanceid%3D%221%22+sessionid%3D%22scifsession_tas-sipnl_crf-tas-dev-b2bua-tas-sipnl-0_1_1_1%22+starttime%3D%222024-03-07+20%3A30%3A25.189%22+stoptime%3D%222024-03-07+20%3A30%3A25.763%22+version%3D%221.0%22%3E%0A+++%3Ccall+call-id%3D%22crf-tas-dev-b2bua-tas-sipnl-0-SIP%238299497007412791717%22+duration%3D%22573%22+mediaType%3D%22UNKNOWN%22+msec%3D%221%22+networkCallType%3D%22ORIGINATING%22+technology%3D%22sip%22%3E%0A++++++%3Ccall_leg+from%3D%22sip%3A16095329681%40msg.lab.t-mobile.com%22+legName%3D%22Uas%22+msec%3D%22350%22+to%3D%22sip%3A19053547887%3Bphone-context%3Dmsg.lab.t-mobile.com%40msg.lab.t-mobile.com%3Buser%3Dphone%22+type%3D%22ORIGINATING%22%3E%0A+++++++++%3Crequest+direction%3D%22incoming%22+messageName%3D%22INVITE%22+msec%3D%22350%22%2F%3E%0A++++++%3C%2Fcall_leg%3E%0A++++++%3Ccall_leg+from%3D%22sip%3A16095329681%40msg.lab.t-mobile.com%22+legName%3D%22Uas%22+msec%3D%22362%22+to%3D%22sip%3A19053547887%3Bphone-context%3Dmsg.lab.t-mobile.com%40msg.lab.t-mobile.com%3Buser%3Dphone%22+type%3D%22ORIGINATING%22%3E%0A+++++++++%3Cresponse+direction%3D%22outgoing%22+msec%3D%22362%22+responseCode%3D%22100%22%2F%3E%0A++++++%3C%2Fcall_leg%3E%0A++++++%3Cscif_event+CallLeg%3D%22sip%3A16095329681%40msg.lab.t-mobile.com%22+Cause%3D%22NONE%22+Contact%3D%22sip%3A19053547887%3Bphone-context%3Dmsg.lab.t-mobile.com%40msg.lab.t-mobile.com%3Buser%3Dphone%3B%22+msec%3D%22372%22+name%3D%22callStart%22+serviceName%3D%22CrfIldTasService%22%2F%3E%0A++++++%3CappLogInviteRcvd+msec%3D%22373%22%3E%0A+++++++++%3Ccall_id+msec%3D%22373%22%3EFA163E3549EF-173e-25165700-14038-65ea24bb-8d8b7%3C%2Fcall_id%3E%0A+++++++++%3Ccalling_party_uri+msec%3D%22373%22%3EP-Asserted-Identity%3A+%26lt%3Bsip%3A%2B16095329681%40msg.lab.t-mobile.com%3Buser%3Dphone%3Bverstat%3DTN-Validation-Passed%26gt%3B%26%2313%3B%0A%3C%2Fcalling_party_uri%3E%0A+++++++++%3Ccalled_party_uri+msec%3D%22373%22%3Esip%3A%2B19053547887%3Bnpdi%3Bphone-context%3Dmsg.lab.t-mobile.com%40msg.lab.t-mobile.com%3Buser%3Dphone%3C%2Fcalled_party_uri%3E%0A+++++++++%3Cims_charging_id+msec%3D%22373%22%3EP-Charging-Vector%3A+icid-value%3Dpmp2.beatf210.sip.lab.t-mobil-1709-843418-869528-680%3Borig-ioi%3D22345%3Bicid-generated-at%3Dpmp2.beatf210.sip.lab.t-mobile.com%26%2313%3B%0A%3C%2Fims_charging_id%3E%0A+++++++++%3Ctermination_type+msec%3D%22373%22%3EILD%3C%2Ftermination_type%3E%0A+++++++++%3Ccli_enabled+msec%3D%22373%22%3EY%3C%2Fcli_enabled%3E%0A+++++++++%3Csource_ip_addr+msec%3D%22373%22%3E10.178.168.201%3C%2Fsource_ip_addr%3E%0A++++++%3C%2FappLogInviteRcvd%3E%0A++++++%3CappLogSipResp+msec%3D%22562%22%3E%0A+++++++++%3Cresp_code+msec%3D%22562%22%3E488%3C%2Fresp_code%3E%0A+++++++++%3Ccrf_resp_reason+msec%3D%22562%22%3ECRF+responded+with+an+error.%3C%2Fcrf_resp_reason%3E%0A++++++%3C%2FappLogSipResp%3E%0A++++++%3Cservice_event+Cause%3D%22SETUP_FAILED%22+msec%3D%22563%22+name%3D%22abortCall%22+serviceName%3D%22CrfIldTasService%22%2F%3E%0A++++++%3Ccall_leg+from%3D%22sip%3A16095329681%40msg.lab.t-mobile.com%22+legName%3D%22caller%22+msec%3D%22573%22+to%3D%22sip%3A19053547887%3Bphone-context%3Dmsg.lab.t-mobile.com%40msg.lab.t-mobile.com%3Buser%3Dphone%22+type%3D%22ORIGINATING%22%3E%0A+++++++++%3Cresponse+direction%3D%22outgoing%22+msec%3D%22573%22+responseCode%3D%22404%22%2F%3E%0A++++++%3C%2Fcall_leg%3E%0A+++%3C%2Fcall%3E%0A%3C%2FCDR%3E%0A`

	// ParseXMLtoStruct("../CDRSample.xml", p)
	ParseXMLtoStruct2(xmlEncodedData, p)
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
	t.Run("Parse appLogSipResp resp_code ", func(t *testing.T) {
		want := []int{400, 180, 200}
		appLogLen := len(p.Call.AppLogSipResp)
		got := make([]int, appLogLen)
		for i := 0; i < appLogLen; i++ {
			got[i] = p.Call.AppLogSipResp[i].RespCode
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
