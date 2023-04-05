package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

// установим значение таймаута по умолчанию (10 секунд)
const defaultTimeout = 10 * time.Second

func main() {
	// определение флагов командной строки
	timeout := flag.Duration("timeout", defaultTimeout, "timeout for connection to server")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet [--timeout duration] host port")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]

	// подключаемся к серверу
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s:%s\n", host, port)

	// запустим горутины для обработки ввода/вывода
	done := make(chan struct{})
	go handleOutput(conn, done)
	go handleInput(conn, done)

	// ждем завершения обработки
	<-done
}

// выводит данные, полученные от сервера, в STDOUT
func handleOutput(conn net.Conn, done chan struct{}) {
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Printf("Failed to read from server: %v\n", err)
	}
	close(done)
}

// отправляет данные из STDIN на сервер
func handleInput(conn net.Conn, done chan struct{}) {
	_, err := io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Printf("Failed to send data to server: %v\n", err)
	}
	conn.Close()
}
