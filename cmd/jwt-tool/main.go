package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tea-jobs/tp-examples/jwt"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Data or signature is missed.")
		return
	}
	mode := strings.TrimSpace(strings.ToLower(os.Args[1]))

	switch mode {
	case "-g":
		generateToken(os.Args[2], os.Args[3])
	case "-p":
		parseToken(os.Args[2], os.Args[3])
	}
}

func generateToken(data string, key string) {
	token, err := jwt.GenerateToken(data, key)
	if err != nil {
		fmt.Printf("Generate a token failure, got: %s\n", err.Error())
		return
	}
	fmt.Printf("Token: %s\n", token)
}

func parseToken(token string, key string) {
	data, err := jwt.ParseToken(token, key)
	if err != nil {
		fmt.Printf("Parse a token failure, got: %s\n", err.Error())
		return
	}
	fmt.Printf("Data: %s\n", data)
}
