package main

import (
	"log"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(conn)

	_, err = conn.Write([]byte("Hello, gopher !\n"))
	if err != nil {
		log.Println(err.Error())
		return
	}

	time.Sleep(time.Second*1)
	return
}
