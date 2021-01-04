package main

import (
	"fmt"
	"github.com/racingmars/go3270"
	"github.com/rs/zerolog/log"
	"io"
	"net"
	"sync"
	"time"
)

func showConnectionError(destinationHost string, destinationPort int32, errormsg error, conn net.Conn) bool {
	fieldValues := make(map[string]string)
	fieldValues["destination"] = destinationHost
	fieldValues["port"] = fmt.Sprintf("%d", destinationPort)
	fieldValues["errormsg"] = errormsg.Error()

	response, err := go3270.HandleScreen(
		connectionErrorScreen,
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
		return false
	}

	return true
}

func readAndFeed(name string, in, out net.Conn, wg *sync.WaitGroup, end, done chan bool) {
	defer func() {
		close(done)
		in.SetReadDeadline(time.Time{})
		log.Debug().Msgf("ending readAndFeed(): %s", name)
		wg.Done()
	}()
	log.Debug().Msgf("starting readAndFeed(): %s", name)
	buffer := make([]byte, 1024)
	finish := false
	for !finish {
		select {
		case <-end:
			log.Debug().Msgf("%s got end signal", name)
			finish = true
		default:
			in.SetReadDeadline(time.Now().Add(time.Second / 2))
			n, err := in.Read(buffer)
			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				continue
			} else if err == io.EOF {
				log.Debug().Msgf("connection closed: %s", name)
				return
			} else if err != nil {
				log.Error().Err(err).Msgf("read error: %s", name)
				return
			}
			log.Trace().Hex("data", buffer[:n]).Msgf("%s read", name)
			if _, err := out.Write(buffer[:n]); err != nil {
				log.Error().Err(err).Msgf("write error: %s", name)
				return
			}
		}
	}
}

func proxy(destinationHost string, destinationPort int32, conn net.Conn) bool {
	target := fmt.Sprintf("%s:%d", destinationHost, destinationPort)

	server, err := net.DialTimeout("tcp", target, 15 * time.Second)

	if err != nil {
		return showConnectionError(destinationHost, destinationPort ,err , conn)
	}
	defer server.Close()

	clientdone := make(chan bool)
	clientend := make(chan bool)
	serverdone := make(chan bool)
	serverend := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go readAndFeed("client", conn, server, &wg, clientend, clientdone)
	go readAndFeed("server", server, conn, &wg, serverend, serverdone)

	select {
	case <-serverdone:
		log.Debug().Msg("got serverdone")
		clientend <- true
	case <-clientdone:
		log.Debug().Msg("got clientdone")
		serverend <- true
	}

	wg.Wait()

	return false
}