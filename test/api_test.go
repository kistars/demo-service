package test

import (
	"fmt"
	"net"
	"os"
	"testing"
)

func TestApi(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("err!")
	} else {
		fmt.Println(hostname)
	}
	addrList, _ := net.LookupIP(hostname)
	for _, addr := range addrList {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println(ipv4)
		}
	}
}
