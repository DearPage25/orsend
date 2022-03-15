package main

import (
	"fmt"
	"net"
)

func (s *server) newClientR(data *DataJson,c *client) {
	fmt.Printf("new client has connected: %s\n", c.conn.RemoteAddr().String())
	
	channelName := data.CHANNEL
	ch, ok := s.channels[channelName]
	if !ok {
		ch = &channel{
			members: make(map[net.Addr]*client),
		}
		s.channels[channelName] = ch
	}
	ch.members[c.conn.RemoteAddr()] = c
}