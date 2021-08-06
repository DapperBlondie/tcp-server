package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
)

type TCPConfig struct {
	CMutex  sync.Mutex
	Counter int32
}

var conf *TCPConfig

func main()  {
	conf = &TCPConfig{
		CMutex:  sync.Mutex{},
		Counter: 0,
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(listener)

	fmt.Println("TCP server is listening on localhost:8080 ...")
	go func() {
		for  {
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err.Error() + "; Occurred in accepting connection.")
				continue
			}

			go handleDialing(conn)
		}
	}()

	<- sigChan
	return
}

func handleDialing(conn net.Conn)  {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(conn)

	str, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(str)

	conf.CMutex.Lock()
	conf.Counter -= 1
	conf.CMutex.Unlock()

	return
}
