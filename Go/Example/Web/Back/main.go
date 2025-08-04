package main

import "net/http"

func runTimeServer(res http.ResponseWriter, req *http.Request) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Test HTTP Server"))
	})
}

func main() {
	http.HandleFunc("/", runTimeServer)

	http.ListenAndServe(":8080", nil)
}
