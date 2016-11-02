package uaparser

import (
	"encoding/json"
	"io/ioutil"
)

type RequestInfo struct {
	UserAgent string
	Browser   string
	OS        string
	Engine    string
	Mobile    string
	Bot       string
}

type LogMessage struct {
	SessionID      int64
	SesStartTime   int
	SesAllTime     int
	ClientIP       string
	BfeIP          string
	Vip            string
	KeepAliveNum   int
	InitialCwnd    int
	RetransRate    int
	Mss            int
	Rtt            int
	SynRtt         int
	RttVar         int
	RTO            int
	ClientInitRwnd int
	ClientMaxRwnd  int
	RequestInfo    []RequestInfo
}

func NewFromJSON(regexJSONFile string) (*Parser, error) {
	data, err := ioutil.ReadFile(regexJSONFile)
	if err != nil {
		return nil, err
	}
	matchIdxNotOk = cDefaultMatchIdxNotOk
	missesTreshold = cDefaultMissesTreshold
	parser, err := NewFromJSONBytes(data)
	if err != nil {
		return nil, err
	}
	return parser, nil
}

func NewFromJSONBytes(data []byte) (*Parser, error) {
	var definitions RegexesDefinitions
	if err := json.Unmarshal(data, &definitions); err != nil {
		return nil, err
	}
	parser := &Parser{definitions, 0, 0, 0, (EOsLookUpMode | EUserAgentLookUpMode | EDeviceLookUpMode), false, false}
	parser.mustCompile()
	return parser, nil
}
