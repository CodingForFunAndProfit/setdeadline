package main

import (
	"crypto/rand"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				return
			}
			go func(c net.Conn) {
				defer c.Close()

				conn.SetDeadline(time.Now().Add(2 * time.Second))
				if err != nil {
					log.Print(err)
					return
				}
				buf := make([]byte, 1<<19) // 512 KB
				for {
					_, err := conn.Read(buf)
					if err != nil {
						log.Print(err)
						break
					}
				}
			}(conn)
		}
	}()

	payload := make([]byte, 1<<20)
	_, err = rand.Read(payload) // generate a random payload
	if err != nil {
		log.Print(err)
	}

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to server.")

	time.Sleep(5 * time.Second)

	_, err = conn.Write(payload)
	if err != nil {
		log.Print(err)
	}

	listener.Close()
}
