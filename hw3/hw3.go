package main

import (
	"time"
)

type Status int

// Status of the cake
const (
	raw Status = iota
	cooked
	garnished
	decorated
)

type Cake struct {
	status Status
	id     int
}

func main() {

	raw_ch := make(chan Cake, 5)

	for i := 1; i <= 5; i++ {
		raw_ch <- Cake{status: raw, id: i}
	}
	close(raw_ch)

	cooked_ch := make(chan Cake, 2)
	garnished_ch := make(chan Cake, 2)
	completed_ch := make(chan Cake, 5)

	go cook(raw_ch, cooked_ch)
	go garnish(cooked_ch, garnished_ch)
	go decorate(garnished_ch, completed_ch)

	for cake := range completed_ch {
		println("Cake", cake.id, "is ready!")
	}
}

func cook(raw_ch chan Cake, cooked_ch chan Cake) {

	for cake := range raw_ch {

		// Critical section
		cake := cake
		time.Sleep(2 * time.Second)
		cake.status = cooked
		println(cake.id, cake.status)
		// End CS

		cooked_ch <- cake
	}
	close(cooked_ch)
	println("Cooked channel closed")
}

func garnish(cooked_ch chan Cake, garnished_ch chan Cake) {

	for cake := range cooked_ch {

		time.Sleep(4 * time.Second)
		cake.status = garnished
		println(cake.id, cake.status)

		garnished_ch <- cake
	}
	close(garnished_ch)
	println("Garnished channel closed")
}

func decorate(garnished_ch chan Cake, completed_ch chan Cake) {

	for cake := range garnished_ch {

		time.Sleep(8 * time.Second)
		cake.status = decorated
		println(cake.id, cake.status)

		completed_ch <- cake
	}
	close(completed_ch)
	println("Completed channel closed")

}
