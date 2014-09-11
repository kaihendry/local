package main

import (
	"./ip"
	"fmt"
	"os"
)

func main() {

	ip, err := ip.Local()

	if err != nil {
		panic(err)
	}

	// local IP: fe80::1 on MACOSX
	fmt.Println("DISCOURAGED local IP:", ip)

	h, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	fmt.Println("ENCOURAGED Hostname:", h)

}
