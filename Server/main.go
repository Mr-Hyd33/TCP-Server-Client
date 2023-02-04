package main

import (
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	log.Println("Starting TCP Server!")
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	handleErrors(err)
	// If an error occurs, we will defer the Listeners handle
	defer listen.Close()

	// Now we'll run an infinite loop to handle all connections
	// that come in, until the server is closed
	for {
		conn, err := listen.Accept()
		handleErrors(err)
		go handleRequests(conn)
	}
}

// Function to handle the Connection Requests
func handleRequests(conn net.Conn) {
	// Make a buffer, size of 1024, to hold the requests
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	handleClientErrors(err)

	// We take the recieved data and write it to a response
	time := time.Now().Format(time.ANSIC)
	log.Printf("Your message: %v. \nRecieved at: %v", string(buffer[:]), time)
	responseString := "This message we will send back to the client"
	conn.Write([]byte(responseString))

	// Finally we will close the Connection
	conn.Close()
}

// Function to handle Client errors [Server will not Exit, Client connection will Close]
func handleClientErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Function to handle Server errors, keep function code clean. [Server will Exit on Error]
func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
