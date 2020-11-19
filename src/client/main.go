package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {

	// Подключаемся к сокету
	conn, _ := net.Dial("tcp", "127.0.0.1:8181")
	for{
		text := "get joke\n"
		// Отправляем в socket
		if _, err := fmt.Fprintf(conn, text); err != nil {
			panic(err.Error())
		}
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
		// sends request once per minute
		time.Sleep(1*time.Minute)
	}
	conn.Close()
}
