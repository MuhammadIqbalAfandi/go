//go:build ignore

package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "", "database password")
	host := flag.String("host", "localhost", "database host")
	port := flag.Int("port", 3306, "database port")

	flag.Parse()

	fmt.Printf("Username: %s\n", *username)
	fmt.Printf("Password: %s\n", *password)
	fmt.Printf("Host: %s\n", *host)
	fmt.Printf("Port: %d\n", *port)
	
		// To run this program, use the following command:
		// go run flag.go -username=admin -password=secret -host=db.example.com -port=5432
	}