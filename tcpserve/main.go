package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr := &net.TCPAddr{
			IP: net.ParseIP("127.0.0.1"),
			Port: 8080,
		}
	l, e := net.ListenTCP("tcp4", addr)
	if e != nil {
		fmt.Println(e)
	}
	/*
	conn, e := l.AcceptTCP()
	if e != nil {
		fmt.Println(e)
	}
	*/

	cons := []net.Conn{}
	go func(cons []net.Conn){
		for {
			c, err := l.Accept()
			if err != nil {
				fmt.Println(err)
			}
			cons = append(cons, c)
		}
	}(cons)
	for {
		for _,con := range cons {
			go func(net.Conn){
				b := []byte{}
				con.Read(b)
			}(con)
			time.Sleep(1 * time.Second)
		}
	}
}
