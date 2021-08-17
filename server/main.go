package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
)
const reply = "HeLlO!\n"
func main() {
	http.HandleFunc("/", kindOfHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func kindOfHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving request")
	tick := time.NewTicker(10 * time.Millisecond)
	defer tick.Stop()
	tock := time.NewTicker(24 * time.Hour)
	defer tock.Stop()
	out:
	for {
		select {
		case <-tock.C:
			break out
		case <-tick.C:
			fmt.Fprint(w, reply)
		default:
			time.Sleep(5 * time.Millisecond)
		}
	}
	fmt.Fprintf(w, "\nWhy are you still here?\n")
}
