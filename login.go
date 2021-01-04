package main

import (
	"github.com/pquerna/otp/totp"
	"github.com/racingmars/go3270"
	"github.com/rs/zerolog/log"
	"net"
	"strings"
)

func validate_mfa(username string, secret string, conn net.Conn) bool {
	fieldValues := make(map[string]string)

	for {
		response, err := go3270.HandleScreen(
			MFAScreen,
			MFAScreenRules,
			fieldValues,
			[]go3270.AID{go3270.AIDEnter},
			[]go3270.AID{go3270.AIDPF3},
			"errormsg",
			8,46,
			conn,
		)
		if err != nil {
			log.Error().Err(err).Msg("Could not deliver MFA screen to client")
			return false
		}

		if response.AID == go3270.AIDPF3 {
			return false
		}
		fieldValues = response.Values

		mfaStr := fieldValues["mfatoken"]
		if totp.Validate(mfaStr, secret) {
			fieldValues["errormsg"] = ""
			log.Info().Msgf("Login for user %s successfully validated using MFA.", username)
			return true
		} else {
			log.Info().Msgf("Invalid MFA token entered for user %s", username)
			fieldValues["errormsg"] = "Invalid OTP entered. Please try again."
		}
	}
}

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
				if val.TOTPKey != "" {
					if ! validate_mfa(username, val.TOTPKey, conn) {
						continue
					}
				}
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