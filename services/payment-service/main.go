package main

import "fmt"

func main() {
	fmt.Println("Payment Gateway API v1")
}

func ProcessPayment(amount float64) string {
	// Simulate payment logic checking it
	if amount <= 0 {
		return "failed"
	}
	return "success"
}

// ff
