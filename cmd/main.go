package main

import (
	"fmt"
	"go-tatar-example/pkg/transport"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	r := transport.Router()
	fmt.Println(http.ListenAndServe(":"+port, r))
}
