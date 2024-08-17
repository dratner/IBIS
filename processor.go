package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dratner/gogpt"
)

func SMSRouter(from string, body string) {
	words := strings.Fields(body)

	cmd, arg := "", ""

	if len(words) > 0 {
		cmd = words[0]
	}
	if len(words) > 1 {
		arg = words[1]
	}

	switch cmd {
	case "on": // this tells the system to forward relevant messages
		log.Printf("User %s signed in", from)
	case "off": // this tells the system to stop forwarding messages
		log.Printf("User %s signed out", from)
	case "status": // replies with your current status (on/off)
		log.Printf("User %s requested status", from)
	case "add": // adds a keyword for filtering
		log.Printf("User %s added the term %s for filtering", from, arg)
	case "remove": // removes a keyword for filtering
		log.Printf("User %s removed the term %s from filtering", from, arg)
	case "all": // overrides filtering and forwards all messages (wildcard)
		log.Printf("User %s requested all terms to be forwarded", from)
	case "keywords": // lists your current keywords
		log.Printf("User %s requested a list of registered keywords", from)
	case "register": // adds a new volunteer with the given number
		log.Printf("User %s registered a new volunteer with number %s", from, arg)
	case "delete": // removes a volunteer with the given number
		log.Printf("User %s deleted a volunteer with number %s", from, arg)
	case "block": // blocks a spammer
		log.Printf("User %s has blocked the caller %s", from, arg)
	case "unblock": // unblocks a spammer
		log.Printf("User %s has unblocked a caller %s", from, arg)
	default:
		log.Printf("Caller %s has sent a message: %s", from, body)
		data, err := ExtractData(body)
		if err != nil {
			log.Printf("could not extract data: %v", err)
		} else {
			log.Printf("extracted data: %+v", data)
		}
		// SendMessage(from, "Thanks from IBIS!")
	}
}

type Location struct {
	City  string `json:"city"`
	State string `json:"state"`
	Zip   string `json:"zipcode"`
}

type MsgData struct {
	Locations []Location `json:"locations"`
}

const prompt = `
Please examine the following text message which relates to a bird or other animal in distress. Do your best to determine which towns or cities in Illinois or Southern Wisconsin this request for assistance refers to. Please reply with a list of each city, state, and the associated zipcode.

Your response must be in the form of a call to the attached function.

##

`

func ExtractData(msg string) (*MsgData, error) {

	key := os.Getenv("OPENAI_KEY")
	if key == "" {
		return nil, fmt.Errorf("openAI key not found")
	}

	query := gogpt.NewGoGPTQuery(key)
	query.Model = "gpt-4o"
	query.Temperature = 0.1
	query.AddMessage(gogpt.ROLE_USER, "", prompt+msg)
	query.AddFunction("is-match", "Call this to provide the data from the message.", MsgData{})

	resp, err := query.Generate()

	if err != nil {
		return nil, err
	}

	reply := resp.Choices[0].Message

	if reply.FunctionCall == nil {
		return nil, fmt.Errorf("non-function response: %s", reply.Content)
	}

	if reply.FunctionCall.Name != "is-match" {
		return nil, fmt.Errorf("unrecognized function response")
	}

	data := new(MsgData)
	err = json.Unmarshal([]byte(reply.FunctionCall.Arguments), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
