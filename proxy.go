package main

import (
	"fmt"
	"github.com/racingmars/go3270"
	"net"
)

func proxy(destinationHost string, destinationPort int32, conn net.Conn) bool {
	fieldValues := make(map[string]string)
	fieldValues["destination"] = destinationHost
	fieldValues["port"] = fmt.Sprintf("%d", destinationPort)

	response, err := go3270.HandleScreen(
		patchingThroughScreen,
		nil,
		fieldValues,
		[]go3270.AID{go3270.AIDEnter},
		[]go3270.AID{go3270.AIDPF3},
		"errormsg",
		0,0,
		conn,
	)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if response.AID == go3270.AIDPF3 {
		return true
	}
	return false
}