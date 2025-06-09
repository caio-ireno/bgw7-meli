package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "My-api-key", "api-key")

	fmt.Printf("My api key is: %s\n", ctx.Value("My-api-key"))

	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	select {
	case <-ctx2.Done():
		fmt.Println("Time Over")
	}
}
