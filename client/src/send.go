package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)


func handleSend(conn net.Conn,  arch string, channel string) {
	var msg string
	path := "./docs/" + arch
	file,err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("sorry, the file don't exist: %s", err)
		defer conn.Close()
	}
	
	if err == nil{
		data := DataJson{"SEND",file, arch, channel}
		json.NewEncoder(conn).Encode(data)
		
		for {
			err = gob.NewDecoder(conn).Decode(&msg)
			if err != nil{
				log.Fatal("sorry, the data isn't convert to json", err)
				break
			}
			if msg != ""{
				fmt.Printf("server: %s\n", msg)	
				break
			}
		}

	}
	
}

