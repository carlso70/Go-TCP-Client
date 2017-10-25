package main

import (
	"net"
	"os"
)

//servAddr := "18.221.200.72:3333"

func main() {
	strEcho := "{ \"userId\": 499380, \"gameId\": 499380, \"inGame\":false }"
	servAddr := "0.0.0.0:3333"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	for {
		_, err = conn.Read(reply)
		if err != nil {
			continue
		}
		println("reply from server=", string(reply))
	}

	defer conn.Close()
}
