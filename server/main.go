package main

import (
	"fmt"
	"net"
)

func main() {
	ls, err := net.Listen("tcp4", ":7000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println("Connection err", err)
			continue
		}
		handler(conn)
	}
	
}

func handler(conn net.Conn) {
	fmt.Println("Connection accepted", conn.RemoteAddr().String())
	for {
		buf := make([]byte, 8)
		_, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read error", err)
			break
		}
		fmt.Printf("message client:%v", string(buf))
		_, err = conn.Write(buf)
		if err != nil {
			fmt.Println("write error", err)
			continue
		}
		
	}
}
