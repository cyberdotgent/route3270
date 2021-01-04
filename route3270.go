package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/racingmars/go3270"
	"net"
	"os"
)

func init() {
	// put the go3270 library in debug mode
	go3270.Debug = os.Stderr
}

var configFile string
func main() {

	parser := argparse.NewParser("route3270", "TN3270 connection router/proxy")

	configFileName := parser.String("c", "config", &argparse.Options{Required: true, Help: "Configuration file to use"})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	configFile = *configFileName

	config := parseConfig(configFile)


	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("LISTENING ON PORT %d FOR CONNECTIONS\n", config.Port)
	fmt.Println("Press Ctrl-C to end server.")
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

// handle is the handler for individual user connections.
func handle(conn net.Conn) {
	defer conn.Close()

	// Always begin new connection by negotiating the telnet options
	go3270.NegotiateTelnet(conn)

	// Handle logging in
	login(conn)


	fmt.Println("Connection closed")
}