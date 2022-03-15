package main

import (
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)

func Commands() {
	conn, err := net.Dial("tcp", ":9000")
	if err  != nil {
		log.Fatal("Error | error tried connection with server: ", err)
		defer conn.Close()
	}

	

	cmd := os.Args[1];
	switch cmd {
		case "receive":
			receiveMode(&conn)
		case "send":
			sendMode(conn)
		default:
			log.Fatal("Accions unknowns | Options: receive or send.")
		
	}
	defer conn.Close()
	
}

func sendMode(conn net.Conn){

	app := &cli.App{}
	app.UseShortOptionHandling =true
	app.Commands  = []*cli.Command{
		{	
			Name:  "send",
			Usage: "send a file mode",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "file", Aliases:[]string{"s"}},
				&cli.StringFlag{Name: "channel", Aliases:[]string{"ch"}},
			},
			
			Action: func (c *cli.Context) error  {
				handleSend(conn, c.String("file"), c.String("channel"))
				return nil
			},
		},
	}
	  
	
	err := app.Run(os.Args)
	if err != nil {
	log.Fatal(err)
	}
}

func receiveMode(conn *net.Conn){
	app := &cli.App{
		Commands: []*cli.Command{
		  {
			Name:        "receive",
			Aliases:     []string{"r"},
			Usage:       "receive a files mode",
			Subcommands: []*cli.Command{
			  {
				Name:  "channel",
				Aliases: []string{"ch"},
				Usage: "channel number",
				Action: func(c *cli.Context) error {
					handleReceive(conn, c.Args().First())
				  	return nil
				},
			  },
			},
		  },
		},
	  }
	
	err := app.Run(os.Args)
	if err != nil {
	log.Fatal(err)
	}
}

