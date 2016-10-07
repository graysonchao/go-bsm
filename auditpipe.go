//build +darwin

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AuditPipeClient struct {
	auditPipe *os.File
	scanner   *bufio.Scanner
}

type AuditPipeMessage struct {
	data string
}

var el = log.New(os.Stderr, "", 0)

func NewAuditPipeClient() *AuditPipeClient {

	f, err := os.Open("/dev/auditpipe")
	if err != nil {
		el.Fatalln("Could not open /dev/auditpipe:", err)
	}

	client := &AuditPipeClient{
		auditPipe: f,
		scanner:   bufio.NewScanner(f),
	}

	return client
}

func (c *AuditPipeClient) Receive() (*AuditPipeMessage, error) {
	c.scanner.Scan()
	if err := c.scanner.Err(); err != nil {
		log.Fatalln("Could not read auditpipe:", err)
	}

	data := c.scanner.Text()
	return &AuditPipeMessage{
		data: data,
	}, nil
}

func main() {
	/*
		c := NewAuditPipeClient()
		for c.scanner.Scan() {
			fmt.Println("Receiving...")
			msg := c.scanner.Text()
			if err := c.scanner.Err(); err != nil {
				el.Fatalln("Receiving message failed:", err)
				os.Exit(1)
			}
			fmt.Println(msg)
		}*/
	f, _ := os.Open("/dev/auditpipe")
	r := bufio.NewReader(f)
	for {
		fmt.Println(r.ReadByte())
	}
}
