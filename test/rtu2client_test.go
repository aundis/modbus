package test

import (
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/goburrow/modbus"
)

func TestRTU2ClientAdvancedUsage(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:5020")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	handler := modbus.NewRTUTCPClientHandler(conn)
	handler.Timeout = 5 * time.Second
	handler.SlaveId = 1
	handler.Logger = log.New(os.Stdout, "tcp: ", log.LstdFlags)

	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(300, 1)
	if err != nil || results == nil {
		t.Fatal(err, results)
	}
	results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
	if err != nil || results == nil {
		t.Fatal(err, results)
	}
	results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})
	if err != nil || results == nil {
		t.Fatal(err, results)
	}
}
