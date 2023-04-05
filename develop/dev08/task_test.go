package main

import (
	"bytes"
	"net"
	"os"
	"strings"
	"testing"
)

func TestShellCommands(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "shell_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"pwd", "pwd\nexit\n", tempDir},
		{"echo", "echo hello world\nexit\n", "hello world"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Chdir(tempDir)
			var output bytes.Buffer
			shell(strings.NewReader(test.input), &output)

			if !strings.Contains(output.String(), test.expected) {
				t.Errorf("Expected output to contain '%s', got '%s'", test.expected, output.String())
			}
		})
	}
}

func TestNetcat(t *testing.T) {
	testServer, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer testServer.Close()

	go func() {
		conn, _ := testServer.Accept()
		defer conn.Close()

		buf := make([]byte, 512)
		n, _ := conn.Read(buf)
		conn.Write(buf[:n])
	}()

	testAddr := testServer.Addr().String()
	input := strings.NewReader("test message\n")
	var output bytes.Buffer

	netcat("tcp", testAddr, input, &output)
	result := output.String()

	if result != "test message\n" {
		t.Errorf("Expected output 'test message\\n', got '%s'", result)
	}
}
