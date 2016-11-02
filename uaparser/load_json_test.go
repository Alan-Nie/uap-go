package uaparser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"bufio"
)

var testJSONParser *Parser

func init() {
	var err error
	testJSONParser, err = NewFromJSON("regexes.json")
	if err != nil {
		log.Fatal(err)
	}
}

func TestLoadJSON(t *testing.T) {
	file, err := os.Open("sample100.json")
	if err != nil {
		t.Error(err)
	}
	reader := bufio.NewReader(file)
	hit, miss := 0, 0
	var missed []string
	for line, err := reader.ReadBytes('\n'); err == nil; {
		var m LogMessage
		json.Unmarshal(line, &m)
		reqs := m.RequestInfo
		for _, v := range reqs {
			agent := testJSONParser.Parse(v.UserAgent)
			if agent != nil && (agent.UserAgent.Family != "Other" || agent.Os.Family != "Other" || agent.Device.Family != "Other") {
				fmt.Printf("Found! %v\n\tUserAgent: %s\n", agent, v.UserAgent)
				hit++
				break
			} else {
				fmt.Printf("Missed! %v\n", v.UserAgent)
				missed = append(missed, v.UserAgent)
				miss++
			}
		}
		line, err = reader.ReadBytes('\n')
	}
}
