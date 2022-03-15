package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)




func main() {
	var data DataJson
	s := newServer()
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("The Server not start: %s", err.Error())
	}
	
	defer listener.Close()
	log.Printf("started server and lisening on PORT: 9000\n")

	for{
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection not accept: ", err.Error())
			continue
		}


		err = json.NewDecoder(conn).Decode(&data)
		if err != nil {
			continue		
		}
		go s.run(&data, &conn)
		
	
	}
}
