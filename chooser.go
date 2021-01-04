package main

import (
	"fmt"
	"github.com/racingmars/go3270"
	"github.com/rs/zerolog/log"
	"net"
	"strings"
)

func chooser(conn net.Conn, username string) bool {
	fieldValues := make(map[string]string)

	for {
		config := parseConfig(configFile)

		for pos, serverName := range config.Users[username].Servers {
			if server, ok := config.Servers[serverName]; ok {
				if len(server.Mnemoric) > 6 {
					fieldValues[fmt.Sprintf("svc%d", pos+1)] = strings.ToUpper(server.Mnemoric)[0:6]
				} else {
					fieldValues[fmt.Sprintf("svc%d", pos+1)] = strings.ToUpper(server.Mnemoric)[0:len(server.Mnemoric)]
				}
				if len(server.Description) > 60 {
					fieldValues[fmt.Sprintf("desc%d", pos+1)] = server.Description[0:60]
				} else {
					fieldValues[fmt.Sprintf("desc%d", pos+1)] = server.Description[0:len(server.Description)]
				}
			}
		}

		response, err := go3270.HandleScreen(
			serverSelectionScreen,
			serverSelectionScreenRules,
			fieldValues,
			[]go3270.AID{go3270.AIDEnter},
			[]go3270.AID{go3270.AIDPF3, go3270.AIDPF12},
			"errormsg",
			20, 15,
			conn,
		)

		if err != nil {
			log.Error().Err(err).Msg("Could not deliver server selection screen to client.")
			return false
		}

		if response.AID == go3270.AIDPF3 {
			// Exit and disconnect
			return false
		}

		if response.AID == go3270.AIDPF12 {
			// Exit and log out
			return true
		}

		fieldValues = response.Values
		serverMnemoricChosen := strings.ToUpper(fieldValues["service"])

		// locate server based on mnemoric
		found := false
		for key, server := range config.Servers {
			if strings.ToUpper(server.Mnemoric) == serverMnemoricChosen {
				// check if user is permitted to be proxied
				allowed := false
				for _, allowedServer := range config.Users[username].Servers {
					if key == allowedServer {
						allowed = true
					}
				}
				if allowed {
					// proxy user
					found = true
					if !proxy(server.Hostname, server.Port, conn) {
						return false
					}
				}
			}
		}
		if ! found {
			fieldValues["errormsg"] = "Please pick a valid service."
		}

	}
}
