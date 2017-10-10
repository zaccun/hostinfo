package main

import "net/http"

func main() {
	http.Handle("/", http.HandlerFunc(hi))
	http.ListenAndServe("localhost:9999", nil)
}

func hi(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello World!\n"))
}
