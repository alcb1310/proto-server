package main

import (
	"fmt"

	"github.com/alcb1310/proto-server/internal/server"
)

var port uint

func main() {
	port = 8080
	srv := server.New(port)

	fmt.Printf("Server listening on port %d\n", port)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
