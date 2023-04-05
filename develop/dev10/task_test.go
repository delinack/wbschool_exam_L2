package main

import (
	"bytes"
	"io"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestHandleOutput(t *testing.T) {
	// Создаем тестовый сервер
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to create test server: %v", err)
	}
	defer ln.Close()

	go func() {
		conn, _ := ln.Accept()
		defer conn.Close()
		_, _ = conn.Write([]byte("Test message from server\n"))
	}()

	// Подключаемся к тестовому серверу
	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}
	defer conn.Close()

	// Перенаправляем STDOUT в буфер
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	done := make(chan struct{})
	go handleOutput(conn, done)

	// Ждем завершения обработки
	<-done

	// Восстанавливаем STDOUT и читаем данные из буфера
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if !strings.Contains(buf.String(), "Test message from server") {
		t.Errorf("handleOutput() failed: expected \"Test message from server\", got \"%s\"", buf.String())
	}
}

func TestHandleInput(t *testing.T) {
	// Создаем тестовый сервер
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to create test server: %v", err)
	}
	defer ln.Close()

	testMessage := "Test message to server\n"

	go func() {
		conn, _ := ln.Accept()
		defer conn.Close()

		buf := make([]byte, len(testMessage))
		_, _ = conn.Read(buf)

		if string(buf) != testMessage {
			t.Errorf("handleInput() failed: expected \"%s\", got \"%s\"", testMessage, string(buf))
		}
	}()

	// Подключаемся к тестовому серверу
	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}
	defer conn.Close()

	// Перенаправляем STDIN из строки
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	_, _ = io.WriteString(w, testMessage)
	w.Close()

	done := make(chan struct{})
	go handleInput(conn, done)

	// Даем горутине handleInput время на обработку
	time.Sleep(1 * time.Second)

	// Восстанавливаем STDIN
	os.Stdin = oldStdin
}
