package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// cd меняет директорию
func cd(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("cd: missing operand")
	}
	return os.Chdir(args[1])
}

// pwd печатает текущую директорию
func pwd() (string, error) {
	return os.Getwd()
}

// echo печатает аргументы
func echo(args []string) string {
	return strings.Join(args[1:], " ")
}

// kill отправляет сигнал SIGTERM указанному идентификатору процесса
func kill(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("kill: missing process ID")
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	return syscall.Kill(pid, syscall.SIGTERM)
}

// ps запускает команду с предоставленными аргументами
func ps(args []string) (string, error) {
	cmd := exec.Command("ps", args[1:]...)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return output.String(), err
}

func shell(input io.Reader, output io.Writer) {
	reader := bufio.NewReader(input)

	for {
		fmt.Fprint(output, "> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		input = strings.TrimSuffix(input, "\n")
		args := strings.Split(input, " ")

		var result string
		var cmdErr error

		switch args[0] {
		case "cd":
			cmdErr = cd(args)
		case "pwd":
			result, cmdErr = pwd()
		case "echo":
			result = echo(args)
		case "kill":
			cmdErr = kill(args)
		case "ps":
			result, cmdErr = ps(args)
		case "exit":
			return
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = output
			cmd.Stderr = os.Stderr
			cmdErr = cmd.Run()
		}

		if cmdErr != nil {
			fmt.Fprintln(os.Stderr, cmdErr)
		} else {
			fmt.Fprint(output, result)
		}
	}
}

func netcat(protocol, addr string, input io.Reader, output io.Writer) {
	var conn net.Conn
	var err error
	switch strings.ToLower(protocol) {
	case "tcp":
		conn, err = net.Dial("tcp", addr)
	case "udp":
		conn, err = net.Dial("udp", addr)
	default:
		fmt.Println("unsupported protocol:", protocol)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		io.Copy(conn, input)
	}()

	io.Copy(output, conn)
}

func main() {
	protocol := flag.String("p", "tcp", "protocol to use (tcp or udp)")
	host := flag.String("h", "", "host to connect to")
	port := flag.Int("P", 0, "port to connect to")
	flag.Parse()

	if *host == "" || *port == 0 {
		fmt.Println("please specify host and port")
		os.Exit(1)
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)

	netcat(*protocol, addr, os.Stdin, os.Stdout)
}
