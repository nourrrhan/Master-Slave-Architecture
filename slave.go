package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Listen for incoming TCP connections
	listener, err := net.Listen("tcp", ":9050")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// Accept incoming connections from the first device
	conn1, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn1.Close()

	// Receive the string from the first device
	buffer := make([]byte, 1024)
	n, err := conn1.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert bytes to string and print the received message
	receivedMessage := string(buffer[:n])
	fmt.Println("Received message:", receivedMessage)

	// Open the file to be sent
	file, err := os.Open("E:/z.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Connect to the remote device
	conn, err := net.Dial("tcp", "192.168.1.140:9050")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send the file size to the remote device
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := make([]byte, 8)
	fileSizeInt := fileInfo.Size()
	for i := 0; i < 8; i++ {
		fileSize[i] = byte(fileSizeInt >> uint(8*i))
	}
	_, err = conn.Write(fileSize)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the file contents to the remote device
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File sent successfully!")
}
