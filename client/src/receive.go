package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)





func handleReceive(conn *net.Conn, channel string) {
	dj := DataJson{
		TYPE:"RECEIVE", 
		CHANNEL: channel,
	}
	
	err := json.NewEncoder(*conn).Encode(dj)

	if err != nil {
		log.Fatal("Error | try send the petition: \n", err)
	}
	fmt.Printf("Lisent in the channel: %s \n", channel)
	dj.lisentServer(conn)
}
