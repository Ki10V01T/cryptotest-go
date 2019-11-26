package simpleproto

import (
	"fmt"
	"net"
	"time"
)


func SendHand(conn net.Conn, b byte){
	bmas := []byte{b}
	conn.Write(bmas)
}

func HandShakeServer(conn net.Conn, srv_ready byte) {
	//defer conn.Close()
	for {
		// read input message
		clientHand := make([]byte, 1)
		n, err := conn.Read(clientHand)
		if n == 0 || err != nil {
			fmt.Println("Read:", err)
			fmt.Println("Disconnected")
			conn.Close()
			break
		}

		if clientHand[0] == 1 {
			println("Client hand = 1" + string(clientHand))
			SendHand(conn, srv_ready)
			break
		}
	}
}

func HandShakeClient(conn net.Conn) {
	//defer conn.Close()
	for {
		SendHand(conn, 1)
		// read input message
		serverHand := make([]byte, 1)
		n, err := conn.Read(serverHand)
		if n == 0 || err != nil {
			fmt.Println("Read:", err)
			fmt.Println("Disconnected")
			break
		}

		if serverHand[0] == 1 {
			println("OK Server hand = 1")
			break
		}
		if serverHand[0] == 0 {
			timer1 := time.NewTimer(2 * time.Second)
			println("Server hand = 0")
			<-timer1.C
			continue
		}
	}
}