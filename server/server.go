package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

const SERVER_PREFIX = "testserver"
const LISTENING_PORT = "6800"


func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:" + LISTENING_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)  // fork subprocess for client
	}
}
 //\w*[^\W*][^$]*
func parseRawMessage(rawMessage string) []string{
	message := strings.SplitN(rawMessage, "$", 2)
	return message
}

func serverResponse(username, message string) string{
	switch message{
	case "@exit":
		return SERVER_PREFIX + "$" + "Disconnected"
	case "@conf":
		return SERVER_PREFIX + "$" + "Conf list: \nSERVER_NAME=" + SERVER_PREFIX + "\nSERVER_PORT=" + LISTENING_PORT
	default:
		return SERVER_PREFIX + "$" + "Got msg | " + message + " | " + " from user: " + username
	}
}

func logger(username, message string) {
	switch message{
	case "@exit":
		fmt.Println("Client: ", username , "| has been disconnected", " | at: ", time.Now().Format(time.Stamp))
	case "@conf":
		fmt.Println("Client: ", username , "| requested a list of settings", " | at: ", time.Now().Format(time.Stamp))
	default:
		fmt.Println("Client: ", username, "| sent message: ", message, " | at: ", time.Now().Format(time.Stamp))
	}
}


// working with client
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// read input message
		input := make([]byte, 1024 * 4)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read:", err)
			fmt.Println("Disconnected")
			break
		}

		clientRawMsg := string(input[0:n])
		clientClearMsg := parseRawMessage(clientRawMsg)

		// server log into terminal
		logger(clientClearMsg[0], clientClearMsg[1])

		// get server response to client
		serverResponseMessage := serverResponse(clientClearMsg[0], clientClearMsg[1])

		// send data to client
		conn.Write([]byte(serverResponseMessage))

		if clientClearMsg[1] == "@exit"{
			conn.Close()
		}
	}
}