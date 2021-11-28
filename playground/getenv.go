package main

import (
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("PORT:", port)
}

// run with:
// PORT=3000 go run getenv.go
