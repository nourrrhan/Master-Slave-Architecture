// package main

// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	// Listen for incoming TCP connections
// 	listener, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer listener.Close()

// 	// Accept incoming connections
// 	conn, err := listener.Accept()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Receive the string from the remote device
// 	buffer := make([]byte, 1024)
// 	n, err := conn.Read(buffer)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Convert bytes to string and print the received message
// 	receivedMessage := string(buffer[:n])
// 	fmt.Println("Received message:", receivedMessage)

// 	//send to the client
// 	if receivedMessage != "" {
// 		// Connect to the remote device
// 		conn, err = net.Dial("tcp", "192.168.1.140:8080")
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		defer conn.Close()

// 		// Send the string to the remote device
// 		chunk1 := "192.168.1.146:9050"
// 		_, err = conn.Write([]byte(chunk1))
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		fmt.Println("String sent successfully!")
// 	} else {
// 		fmt.Println("wrong handshake")
// 	}
// }
