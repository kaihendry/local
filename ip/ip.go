package ip

import (
	"net"
)

func Local() (string, error) {

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
			// fmt.Println("Debug", addr)
			switch v := addr.(type) {
			case *net.IPNet:
				if !v.IP.IsLoopback() {
					return v.IP.String(), nil
				}
			case *net.IPAddr:
				// fmt.Println("Debug", v.IP)
				if !v.IP.IsLoopback() && v.IP.String() != "0.0.0.0" {
					return v.IP.String(), nil
				}
		
				// fmt.Printf("A %v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())
			}
		}
	}

	return "0.0.0.0", nil
}
