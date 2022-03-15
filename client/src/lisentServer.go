package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"net"
)

type DataJson struct {
	TYPE	string 	`json:"TYPE"`
	FILE 	[]byte 	`json:"FILE"`
	NAME	string 	`json:"NAME"`
	CHANNEL string 	`json:"CHANNEL"`
}

func (dj *DataJson) lisentServer(conn *net.Conn)  {
	for {
		err := gob.NewDecoder(*conn).Decode(&dj)
		if err != nil{
			fmt.Println("The server is down!")
			return
		} 
		fmt.Println("You have a new file")
		saveFile(dj, *conn)
	
	}
	
}

func saveFile(dj *DataJson, conn net.Conn ) {
	
	path := conn.LocalAddr().String() + "-" + dj.NAME 
	
	err := ioutil.WriteFile("./file/"+ path, dj.FILE, 0666)
	
	if err != nil {
		fmt.Println("Error | tried saved the file ", err)
	}
	
}


