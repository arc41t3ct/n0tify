package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func listen(es *eventServer) {
	hypr := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	addr, aErr := net.ResolveUnixAddr("unix", fmt.Sprintf("/tmp/hypr/%s/.socket2.sock", hypr))
	if aErr != nil {
		log.Fatal(aErr)
	}
	conn, cErr := net.DialUnix("unix", nil, addr)
	if cErr != nil {
		log.Fatal(cErr)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		data, dErr := reader.ReadString('\n')
		if dErr != nil {
			log.Fatal(dErr)
		}
		data = strings.ReplaceAll(data, "\n", "")
		es.dispatch(data)
	}
}
