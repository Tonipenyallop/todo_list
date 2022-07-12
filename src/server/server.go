package main

import (
	"fmt"
	"log"

	"net"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "welcome to the golang server!s")
	// })
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen PORT with 9000 %v", err)
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Server(lis); err != nil {
		log.Fatalf("Failed to listen PORT with 9000 %v", err)
	}
	http.HandleFunc("/hello", helloHandler)
	// PORT := 8080
	// fmt.Printf("Listening port with %d", PORT)
	fmt.Printf("Listening port with 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is unsupported", http.StatusNotFound)
		return
	}
	toni := "I am the best of all tOime!"
	fmt.Fprintf(w, toni)
}
