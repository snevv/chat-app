package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var (
	messages = make(chan string)
	entering = make(chan *Client)
	leaving  = make(chan string)
)

type Client struct {
	Addr    string
	Name    string
	MsgChan chan string
	Active  bool
}

func boxedMessage(msg string) string {
	border := strings.Repeat("-", len(msg)+4)
	return fmt.Sprintf("%s\n| %s |\n%s", border, msg, border)
}

func broadcaster() {
	clients := map[string]*Client{}
	for {
		select {
		case msg := <-messages:
			for _, cli := range clients {
				select {
				case cli.MsgChan <- msg:
				default:
				}
			}
		case cli := <-entering:
			clients[cli.Addr] = cli
			fmt.Println("Client", cli.Addr, "has entered...")
		case cli := <-leaving:
			fmt.Println("Client", cli, "has left...")
			close(clients[cli].MsgChan)
			delete(clients, cli)
		}
	}
}

func handleConn(conn net.Conn) {
	timeout := time.NewTimer(5 * time.Minute)
	done := make(chan struct{})

	input := bufio.NewScanner(conn)
	var user string
	fmt.Fprint(conn, "Enter your username: ")
	if input.Scan() {
		user = input.Text()
	}

	addr := conn.RemoteAddr().String()
	msgChan := make(chan string)
	client := &Client{
		Addr:    addr,
		Name:    user,
		MsgChan: msgChan,
		Active:  true,
	}
	messages <- boxedMessage(user + " has entered...")
	entering <- client

	go func() {
		<-timeout.C
		close(done)
		conn.Close()
	}()

	go func(conn net.Conn, ch <-chan string) {
		for msg := range ch {
			fmt.Fprintln(conn, msg)
		}
	}(conn, msgChan)

	for input.Scan() {
		messages <- fmt.Sprintf("[%s]: %s", user, input.Text())

		if !timeout.Stop() {
			<-timeout.C
		}
		timeout.Reset(5 * time.Minute)
	}

	leaving <- addr
	messages <- boxedMessage(user + " has left...")

	select {
	case <-done:
	default:
		conn.Close()
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
