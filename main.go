package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"slack_bot/secrets"
	"slack_bot/slack_structs"
)

func main() {
	http.HandleFunc("/bot", botHandler) // each request calls staticHandler
	log.Fatal(http.ListenAndServe(":8500", nil))
}

func botHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var eventPacket slack_structs.EventWrapper
	if err := json.Unmarshal(body, &eventPacket); err != nil {
		fmt.Fprintf(w, "JSON unmarshalling failed: %s", err)
	}

	// constructing message to return
	messageWrapper := messageBuilder(&eventPacket)
	marshaledMessage, err := json.Marshal(messageWrapper)

	if err != nil {
		fmt.Fprintf(w, "JSON marshalling failed: %s", err)
	}

	// writing message
	fmt.Fprint(w, string(marshaledMessage))

	// writing OK status
	w.WriteHeader(http.StatusOK)
}

func messageBuilder(eventPacket *slack_structs.EventWrapper) slack_structs.MessageWrapper {
	// simple reflection message builder to echo out message in which bot
	// was mentioned in the channel it was mentioned
	message := slack_structs.Message{Text: eventPacket.Event.Text, Ts: eventPacket.Event.Ts}
	messageWrapper := slack_structs.MessageWrapper{
		Token: secrets.BotToken,
		Ok: true,
		Ts: eventPacket.Event.Ts,
		Message: message,
		Channel: eventPacket.Event.Channel,
	}

	return messageWrapper
}