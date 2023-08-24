package main

import "sync"

type Status int

const (
	Raw Status = iota
	Cooked
	Garnished
	Decorated
)

type Cake struct {
	Status Status
}

func cook(cake Cake) {

	cake.Status = Cooked

}

func garnish() {

	cake.Status = Garnished

}

func decorate() {

	cake.Status = Decorated

}



func main() {

	raw_ch := make(chan Cake, 5)
	cook_ch := make(chan Cake, 1)
	garnish_ch := make(chan Cake, 1)
	decorate_ch := make(chan Cake, 1)

	cooked_ch := make(chan Cake, 2) 
	garnished_ch := make(chan Cake, 2)

	var wg sync.WaitGroup

	wg.Add(1)

	go cook(cake)
	go garnish()
	go decorate()

	wg.Wait()

	
	

}
