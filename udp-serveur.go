package main

import (
	"log"
	"net"
	"fmt"
	"os"
)

func main() {
	port := os.Args[1]
	
	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", ":" + port)
	
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		fmt.Print(addr, ": ", string(buf[:n]))
	}

}
