package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp4", ":7000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go func ()  {
		for {
			buf := make([]byte, 8)
			_, err := conn.Read(buf[:])
			if err != nil {
				fmt.Println("read error", err)
				conn.Close()
				break
			}
			fmt.Printf("message SERVER:%v\n", string(buf))
			
		}
	}()
	
	for i := 0; i < 10; i++ {
		_, err = conn.Write([]byte("Hello!!!"))
		if err != nil {
			fmt.Println("write error", err)
			continue
		}
	}

	for{
		select{

		}

	}

}
