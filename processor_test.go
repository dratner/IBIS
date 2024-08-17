package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExtractData(t *testing.T) {

	type EDTest struct {
		msg   string
		city  string
		state string
	}

	edTest := []EDTest{
		{
			msg:   "Special situation in Evanston. Ladder needed to get to secord story of building. Can anyone help?",
			city:  "evanston",
			state: "illinois",
		},
		{
			msg:   "Injured peregrine falcon in yard near Touhy & Harlem. Who can rescue? Contained bird in Monee, who can pick up? Text STOP to quit.",
			city:  "chicago",
			state: "il",
		},
	}

	for _, tst := range edTest {

		data, err := ExtractData(tst.msg)

		if err != nil {
			t.Errorf("function ExtractData() has error %v", err)
		} else {
			t.Logf("data is %+v", data)
		}

		success := false
		result := ""

		for _, loc := range data.Locations {
			city, state := strings.ToLower(loc.City), strings.ToLower(loc.State)
			result = fmt.Sprintf("expected: %s, %s got: %s, %s", tst.city, tst.state[:2], city, state[:2])
			if (city == tst.city) && (state[:2] == tst.state[:2]) {
				success = true
				break
			}
		}

		if success {
			t.Logf("test message correctly identified: %s", result)
		} else {
			t.Errorf("test message wrongly identified: %s", result)
		}
	}
}
