package main
import (
	"fmt"
	"net"
	"os"
	"strings"
)

const CLIENT_PREFIX = "testuser"
const SERVER_PORT = "6800"

func parseRawMessage(rawMessage string) []string{
	message := strings.SplitN(rawMessage, "$", 2)
	return message
}

func clientResponse(message string) string{
	return CLIENT_PREFIX + "$" + message
}


func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:" + SERVER_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for{
		var message string
		fmt.Print("Message" + " > ")
		_, err := fmt.Scanln(&message)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}
		// send msg for server
		readyData := clientResponse(message)
		if n, err := conn.Write([]byte(readyData));
			n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		// get response
		fmt.Print("Response < ")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err !=nil{ break}
		serverRawMsg := string(buff[0:n])
		serverClearMsg := parseRawMessage(serverRawMsg)
		//fmt.Print(string(buff[0:n]))
		fmt.Print(serverClearMsg[1])
		fmt.Println()

		if message == "@exit" {
			conn.Close()
			os.Exit(0)
		}
	}
}