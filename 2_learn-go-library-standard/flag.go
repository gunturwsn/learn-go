package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "root", "database password")
	port := flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("host", *host)
	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("port", *port)

	// running => go run flag.go
	// running => go run flag.go -username=Budi -password=123 -host=123.123.123 -port=8088
}
