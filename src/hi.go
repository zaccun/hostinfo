package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//Set listening port
	port := 9898
	if len(os.Args) > 1 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	//Find default output IP and hostname
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	//Register handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(hostname + " " + localAddr.IP.String() + "\n"))
	})
	//Serve
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(port), nil))
}

