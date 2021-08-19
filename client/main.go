package main

import (
	"fmt"
	"flag"
	"net"
//	"net/http"
	"time"
	"sync"
//	"runtime"
//	"math/rand"
)
type config struct {
	mu sync.Mutex
	cons []net.Conn
	bufferSize *int
	maxSock *int
}

var conf config

func init() {
	conf.bufferSize = flag.Int("b", 828375040, "Size of the buffer you are trying to fill in bytes, default: 828375040")
	conf.maxSock = flag.Int("s", 50000, "Maximum number of sockets to create, default: 50000")
}
func main() {

	flag.Parse()
//	cons := []net.Conn{}
//	count := 0
	errCount := 0

	// This is the max tcp_wmem value on an m5.large
	bs := 1024 * 1024 * 16

	// we need buffMax 1Mi buffers to fill the queue
	buffMax := *conf.bufferSize / bs
	max := *conf.maxSock

	if buffMax < max {
		max = buffMax
	}

	fmt.Println("max: ", max, " BuffMax: ", buffMax)
	for i := 0; i < max; i++ {
		go func(){
			conn, err := net.Dial("tcp", "127.0.0.1:8080")
			if err != nil {
				errCount++
				return
			}
			if tcpC, ok := conn.(*net.TCPConn); ok {
				if err := tcpC.SetWriteBuffer(16777216); err != nil { errCount++; return }
			}

			conf.mu.Lock()
			conf.cons = append(conf.cons, conn)
			conf.mu.Unlock()
		}()
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("cons: %d\t errs: %d\n", len(conf.cons), errCount)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Writing ", len(conf.cons), " conns")
	for i,conn := range conf.cons {
		fmt.Println("writing ", i )
		j := i
		go func(i int,  conn net.Conn){
			s := time.Now()
			_, e := conn.Write(make([]byte, bs)); if e != nil{ errCount++; fmt.Println("failed write") }
			end := time.Now()
			fmt.Printf("Write for %d took %d\n", j, end.Sub(s).Seconds())
		}(j, conn)
	}
	fmt.Printf("cons: %d\t errs: %d\n", len(conf.cons), errCount)
	time.Sleep(24 * time.Hour)
/*
	count := 0
	errCount := 0
	for {
		for i := 0; i < 100; i++ {
			go func(){
				r, err := http.Get("http://127.0.0.1:8080")
				defer r.Body.Close()
				if err != nil { errCount++ }
			}()
			count++
		}
		time.Sleep(1 * time.Second)
		fmt.Println("count: ", count - errCount, " Errors: ", errCount)
	}
*/
}
