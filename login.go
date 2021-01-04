package main

import (
	"github.com/racingmars/go3270"
	"github.com/rs/zerolog/log"
	"net"
	"strings"
)

func login(conn net.Conn) {
	config := parseConfig(configFile)

	fieldValues := make(map[string]string)
	for {

		// ask for the user to log in
		response, err := go3270.HandleScreen(
			loginScreen,
			loginScreenRules,
			fieldValues,
			[]go3270.AID{go3270.AIDEnter},
			[]go3270.AID{go3270.AIDPF3},
			"errormsg",
			6,16,
			conn,
		)
		if err != nil {
			log.Error().Err(err).Msg("Could not deliver login screen to client")
			return
		}

		if response.AID == go3270.AIDPF3 {
			// Exit
			break
		}

		fieldValues = response.Values
		username := strings.TrimSpace(fieldValues["username"])
		password := strings.TrimSpace(fieldValues["password"])

		fieldValues["errormsg"] = "Invalid username or password"
		if val, ok := config.Users[username]; ok {
			if password == val.Password {
				fieldValues["errormsg"] = ""
				log.Info().Msgf("Successful login for user %s from host %s", username , conn.RemoteAddr())
				if chooser(conn, username) {
					continue
				} else {
					log.Info().Msgf("Session ended for user %s", username)
					break
				}
			}
		}
		if username != "" {
			log.Warn().Msgf("Invalid login for user %s from host %s", username, conn.RemoteAddr())
		}

		continue
	}
}