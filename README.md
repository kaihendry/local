Finding the address for local httpd service URL sharing.

	go get github.com/kaihendry/local/ip

# Findings

## MacOSX

	DISCOURAGED local IP: fe80::1
	ENCOURAGED Hostname: Ruths-MBP.local

local IP is a fail, but hostname is a winner!

## (Arch) Linux

	DISCOURAGED local IP: 192.168.88.165
	ENCOURAGED Hostname: x220

local IP is workable, but hostname is missing `.local` suffix

As x220.local, it's fine locally and from other machines.

## Windows

	DISCOURAGED local IP: 192.168.88.137
	ENCOURAGED Hostname: STARCRAFT-PC

local IP is workable, but hostname is missing .local suffix

As STARCRAFT-PC.local is doesn't work on local Windows machine or from other machines on the network.
