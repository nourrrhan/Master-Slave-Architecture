package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Connect to the master
	conn, err := net.Dial("tcp", "192.168.183.254:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send string to the master
	message := "Hi Master"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected Successfully to Master!")

	//---------------------------

	// Recieve slaves information (ip + port number)
	// Listen for incoming TCP connection
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// Accept incoming connection
	conn1, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer conn1.Close()

	// Receive slaves ips and port numbers from master
	buffer := make([]byte, 1024)
	n, err := conn1.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert bytes to string and print the received message
	receivedMessage := string(buffer[:n])
	fmt.Println("Received ips:", receivedMessage)

	// Split slaves ips and port numbers in array and store the number of devices
	addresses := strings.Split(receivedMessage, ",")
	length := len(addresses)

	//----------------------------------------------------

	// Connect to the slaves and recieve files
	for i := 0; i < length; i++ {

		// Start connection
		conn2, err := net.Dial("tcp", addresses[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn2.Close()

		// Send the string to the slave i
		message1 := "Hi Slave"
		_, err = conn2.Write([]byte(message1))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Connected Successfully to Slave!")

		//--------------------------------------------
		// Store each slave port number and start receiving files from slaves
		slaveInfo := addresses[i]
		portNum := slaveInfo[len(slaveInfo)-5:]

		// Listen for incoming TCP connection
		listen, err := net.Listen("tcp", portNum)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer listen.Close()

		// Accept incoming connection
		conn4, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn4.Close()

		// Receive the file size
		fileSizeBytes := make([]byte, 8)
		_, err = io.ReadFull(conn4, fileSizeBytes)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileSizeInt := int64(0)
		for i := uint(0); i < 8; i++ {
			fileSizeInt |= int64(fileSizeBytes[i]) << (8 * i)
		}

		// Receive the file
		// Create a new file to store the received data
		fileName := "received-file" + strconv.Itoa(i+1) + ".txt"
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Receive the file contents and write them to the file
		_, err = io.CopyN(file, conn4, fileSizeInt)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("File Received Successfully!")

	}

}
