package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pair struct {
	name  string
	price chan float64
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	wg := sync.WaitGroup{}

	pairs := []Pair{
		{"EUR/USD", make(chan float64, 1)},
		{"GBP/USD", make(chan float64, 1)},
		{"JPY/USD", make(chan float64, 1)},
	}

	go simulateMarketData(&pairs, &wg)
	go selectPair(&pairs, &wg)

	time.Sleep(1 * time.Minute)
}

func selectPair(pairs *[]Pair, wg *sync.WaitGroup) {
	for {

		select {
		case price := <-(*pairs)[0].price:
			if price > 1.2 {
				formattedPrice := fmt.Sprintf("%.5f", price)
				println(time.Now().Format("15:04:05"), "EUR/USD:", formattedPrice, "- Buying...")
				time.Sleep(4 * time.Second)
				println(time.Now().Format("15:04:05"), "EUR/USD", "- Bought!")
			}
		case price := <-(*pairs)[1].price:
			if price < 1.35 {
				formattedPrice := fmt.Sprintf("%.5f", price)
				println(time.Now().Format("15:04:05"), "GBP/USD:", formattedPrice, "- Buying...")
				time.Sleep(3 * time.Second)
				println(time.Now().Format("15:04:05"), "GBP/USD", "- Bought!")
			}
		case price := <-(*pairs)[2].price:
			if price < 0.0085 {
				formattedPrice := fmt.Sprintf("%.5f", price)
				println(time.Now().Format("15:04:05"), "JPY/USD:", formattedPrice, "- Buying...")
				time.Sleep(3 * time.Second)
				println(time.Now().Format("15:04:05"), "JPY/USD", "- Bought!")
			}

		default:
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func simulateMarketData(pairs *[]Pair, wg *sync.WaitGroup) {
	for {

		randPrice := rand.Float64()*0.5 + 1.0

		// EUR/USD
		(*pairs)[0].price <- randPrice
		// GBP/USD
		(*pairs)[1].price <- randPrice

		randPrice = rand.Float64()*0.003 + 0.006

		// JPY/USD
		(*pairs)[2].price <- randPrice

		time.Sleep(1 * time.Second)
	}
}
