package main

import (
	"log"
	"net"
)

type DataJson struct {
	TYPE	string 	`json:"TYPE"`
	FILE 	[]byte 	`json:"FILE"`
	NAME	string 	`json:"NAME"`
	CHANNEL string 	`json:"CHANNEL"`
}

type server struct {
	channels map[string]*channel
	
}
type client struct {
	conn net.Conn
}



func newServer() *server{
	return &server{
		make(map[string]*channel),
	}
} 
func (s *server) run(data *DataJson, conn *net.Conn) {
	c:= &client{
		conn:  *conn,
	}
	switch data.TYPE {
		case "RECEIVE":
			go s.newClientR(data, c)
		case "SEND":
			s.newClientS(data, c)
			
		default:
			log.Fatal("Error | DON'T select option:", data)
	}	
}




