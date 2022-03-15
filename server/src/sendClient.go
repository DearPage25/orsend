package main

import (
	"encoding/gob"
	"fmt"
)


func (s *server) newClientS(data *DataJson, c *client) {
	channelName := data.CHANNEL
	ch, ok := s.channels[channelName]
	if !ok {
		fmt.Println("Sorry don't exist this channel: ", channelName)
		gob.NewEncoder(c.conn).Encode("Sorry don't exist this channel: " + channelName + "\n")
	}else {
		ch.broadcast(data)
		gob.NewEncoder(c.conn).Encode("Great! Your file was sent.")
	}


}
