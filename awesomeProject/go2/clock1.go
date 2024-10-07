package go2

import (
	"io"
	"log"
	"net"
	"time"
)

func Clock1() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // Например, обрыв соединения
			continue
		}
		go handleConn(conn) // Обработка единственного подключения
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05"))
		if err != nil {
			return // Например, отключение клиента
		}
		time.Sleep(1 * time.Second)
	}
}
