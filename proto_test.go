package main

import (
	"fmt"
	"testing"
)

func TestProtocol(t *testing.T) {

	raw := "*3\r\n$3\r\nSET\r\n$3\r\nKEY\r\n$3\r\nFOO\r\n"

	cmd, err := ParseCommand(raw)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cmd)
}
