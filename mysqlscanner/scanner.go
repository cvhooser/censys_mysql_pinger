package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"unicode"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments.")
		return
	}
	url := args[0] + ":" + args[1]

	conn, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Println("failed:", err)
		log.Fatal("No instance of Mysql detected")
		os.Exit(1)
	}

	info := pingInstance(conn)
	defer conn.Close()

	// The version isn't static so I have to calculate the offset
	if len(info) > 6 {
		version := ""
		offset := 0
		for _, char := range string(info[5:]) {
			if !unicode.IsDigit(char) && char != '.' {
				break
			}
			offset++
			version += string(char)
		}
		offsetBottom := offset + 50
		offsetTop := offsetBottom + 21
		
		// testing offset calculation
		// fmt.Printf("Bottom: %v, Top: %v", offsetBottom, offsetTop)
		// fmt.Printf(info[offsetBottom:offsetTop]);
		if len(info) > offsetTop && info[offsetBottom:offsetTop] == "mysql_native_password" {
			fmt.Printf("Mysql is running version %s at %s\n", version, url)
		}
	} else {
		fmt.Printf("No instance of Mysql detected\n")
	}

}

func pingInstance(conn net.Conn) string {

	// mysql authentication request encoded with root:root 
	// I used the golang mysql driver and captured 
	// the packet data with ngrep
	_, err := conn.Write([]byte{0x55, 0x00, 0x00, 0x01, 0x8d, 0xa2, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x72, 0x6f, 0x6f, 0x74, 0x00, 0x14, 0x25, 0x50, 0x5c, 0x86, 0xc2, 0xdc, 0x52, 0xa8, 0x16, 0xbc, 0xc8, 0x45, 0xc7, 0x0a, 0x92, 0x53, 0x28, 0x33, 0xc1, 0xed, 0x74, 0x65, 0x73, 0x74, 0x00, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x00})
	if err != nil {
		fmt.Println("failed:", err)
		fmt.Printf("No instance of mysql detected\n")
		os.Exit(1)
	}

	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("failed:", err)
		fmt.Printf("No instance of mysql detected\n")
		os.Exit(1)
	}
	// debugging purposes
	// for pos, char := range string(buf) {
	//   fmt.Printf("character %c starts at byte position %d\n", char, pos)
	// }
	// fmt.Printf(string(buf)+ "\n")
	return string(buf)
}
