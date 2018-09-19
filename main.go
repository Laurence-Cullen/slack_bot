package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"slack_bot/slack_structs"
)

func main() {
	http.HandleFunc("/bot", botHandler) // each request calls staticHandler
	log.Fatal(http.ListenAndServe(":8500", nil))
}

func botHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var eventPacket slack_structs.EventPacket
	if err := json.Unmarshal(body, &eventPacket); err != nil{
		fmt.Fprintf(w, "JSON unmarshalling failed: %s", err)
	}

	fmt.Fprintf(w, "event packet: %v", eventPacket)

	//fmt.Fprint(w, string(body))
}