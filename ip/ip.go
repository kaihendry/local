package ip

import (
	"errors"
	"fmt"
	"net"
)

func ExternalIP() (string, error) {

	ifaces, err := net.Interfaces()

	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()

		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			fmt.Println(addr)
			switch v := addr.(type) {
			case *net.IPNet:
				if !v.IP.IsLoopback() {
					return v.IP.String(), nil
				}
			case *net.IPAddr:
				fmt.Printf("%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())
			}
		}
	}

	return "", errors.New("are you connected to the network?")
}
