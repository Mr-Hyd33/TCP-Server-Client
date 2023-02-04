package main

import (
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Println("Unable to resolve the given address\n\tERROR: " + err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		log.Printf("Unable to make a connection to the Server.\n\tERROR: %v", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte("This message is to the Server from the Client"))
	if err != nil {
		log.Printf("Could not write the Message data.\n\tERROR: %v", err.Error())
		os.Exit(1)
	}

	// We will create a buffer here to recieve data.
	recievedBuffer := make([]byte, 1024)
	_, err = conn.Read(recievedBuffer)
	if err != nil {
		log.Printf("Could not recieve data sent from the server.\n\tERROR: %v", err.Error())
		os.Exit(1)
	}

	log.Printf("Recieved Message Content: \n\t%v", string(recievedBuffer))

	// Finally once the message has been sent and recieved, we close the connection
	conn.Close()
}
