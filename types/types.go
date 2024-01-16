package types

import "encoding/xml"

type Person struct {
	XMLName xml.Name `xml:"person"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name"`
}

type CDR struct {
	XMLName xml.Name `xml:"CDR"`
	// DebugNamespace  string   `xml:"xmlns:debug,attr"`
	// LogNamespace    string   `xml:"xmlns:log,attr"`
	ContextVersion  string `xml:"contextversion,attr"`
	Host            string `xml:"host,attr"`
	Level           string `xml:"level,attr"`
	ScifServiceName string `xml:"scifservicename,attr"`
	SeeInstanceID   string `xml:"seeinstanceid,attr"`
	SessionID       string `xml:"sessionid,attr"`
	StartTime       string `xml:"starttime,attr"`
	StopTime        string `xml:"stoptime,attr"`
	Version         string `xml:"version,attr"`
	Call            Call   `xml:"call"`
}

type Call struct {
	CallID           string            `xml:"call-id,attr"`
	Duration         int               `xml:"duration,attr"`
	MediaType        string            `xml:"mediaType,attr"`
	Msec             int               `xml:"msec,attr"`
	NetworkCallType  string            `xml:"networkCallType,attr"`
	Technology       string            `xml:"technology,attr"`
	AppLogInviteRcvd AppLogInviteRcvd  `xml:"appLogInviteRcvd"`
	AppLogRoute      AppLogRoute       `xml:"appLogRoute"`
	AppLogInviteFwd  []AppLogInviteFwd `xml:"appLogInviteFwd"`
	AppLogSipResp    []AppLogSipResp   `xml:"appLogSipResp"`
}

type AppLogInviteRcvd struct {
	Msec            int    `xml:"msec,attr"`
	CallID          string `xml:"call_id"`
	CallingPartyURI string `xml:"calling_party_uri"`
	CalledPartyURI  string `xml:"called_party_uri"`
	TerminationType string `xml:"termination_type"`
	CliEnabled      string `xml:"cli_enabled"`
	SourceIPAddr    string `xml:"source_ip_addr"`
}

type AppLogRoute struct {
	Msec     int    `xml:"msec,attr"`
	DialCode string `xml:"dial_code"`
}

type AppLogInviteFwd struct {
	Msec       int    `xml:"msec,attr"`
	Contract   string `xml:"contract"`
	FQDN       string `xml:"fqdn"`
	Tier       string `xml:"tier"`
	TermIPAddr string `xml:"term_ip_addr"`
	Rank       int    `xml:"rank"`
}

type AppLogSipResp struct {
	Msec         int    `xml:"msec,attr"`
	RespCode     int    `xml:"resp_code"`
	ReasonHeader string `xml:"reason_header"`
}
