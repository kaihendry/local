package main

import (
	"./ip"
	"fmt"
)

func main() {

	ip, err := ip.ExternalIP()

	if err != nil {
		panic(err)
	}

	fmt.Println("local IP:", ip)

}
