package main

import (
	"log"
	"strings"
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
		SendMessage(from, "Thanks from IBIS!")
	}
}
