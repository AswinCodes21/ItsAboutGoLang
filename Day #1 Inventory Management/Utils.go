package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateID() func() string {
	rand.Seed(time.Now().UnixNano())
	return func() string {
		return fmt.Sprintf("P%03d", rand.Intn(1000))
	}
}

func CalculateTotal(prices ...float64) float64 {
	total := 0.0
	for _, price := range prices {
		total += price
	}
	return total
}
