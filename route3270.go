package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/racingmars/go3270"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net"
	"os"
	"time"
)


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

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("Starting Route/3270.")

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Error().Msgf("%s", err)
		log.Error().Err(err).Msg("Could not start listening on the specified port. Quitting.")
		return
	}

	log.Info().Msgf("Listening on port %d for connections.", config.Port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Error().Err(err).Msg("Could not accept the connection.")
			continue
		}
		go handle(conn)
	}
}

// handle is the handler for individual user connections.
func handle(conn net.Conn) {
	defer conn.Close()

	log.Info().Msgf("Accepted connection from %s", conn.RemoteAddr())
	// Always begin new connection by negotiating the telnet options
	go3270.NegotiateTelnet(conn)

	// Handle logging in
	login(conn)

}
