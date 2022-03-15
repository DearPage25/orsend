package main

import (
	"encoding/gob"
	"net"
)

type channel struct {
	members map[net.Addr]*client
}

func (c *channel) broadcast(data *DataJson){
	for _, members:= range c.members{
		gob.NewEncoder(members.conn).Encode(data)
	}
}