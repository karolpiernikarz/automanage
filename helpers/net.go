package helpers

import (
	"fmt"
	"net"
)

func GetFreePort() (int, error) {
	// get a free port
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer func(l *net.TCPListener) {
		err := l.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(l)
	return l.Addr().(*net.TCPAddr).Port, nil
}
