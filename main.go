package main

import (
	"./ip"
	"fmt"
)

func main() {

	ip, err := ip.Local()

	if err != nil {
		panic(err)
	}

	fmt.Println("local IP:", ip)

}
