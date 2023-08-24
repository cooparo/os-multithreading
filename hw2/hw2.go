package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Customer struct {
	name string
}

type Vehicle struct {
	v_type string
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	ch := make(chan Vehicle)
	wg := sync.WaitGroup{}

	customers := []Customer{
		{"Cliente 1"},
		{"Cliente 2"},
		{"Cliente 3"},
		{"Cliente 4"},
		{"Cliente 5"},
		{"Cliente 6"},
		{"Cliente 7"},
		{"Cliente 8"},
		{"Cliente 9"},
		{"Cliente 10"},
	}

	for _, customer := range customers {
		go rent(customer, ch, &wg)
	}

	var rented []Vehicle

	for i := 0; i < len(customers); i++ {
		vehicle := <-ch
		rented = append(rented, vehicle)
	}

	wg.Wait()
	close(ch)

	print(rented)
}

func rent(customer Customer, ch chan Vehicle, wg *sync.WaitGroup) Vehicle {
	defer wg.Done()

	vehicles := []Vehicle{
		{"Berlina"},
		{"SUV"},
		{"Station Wagon"},
	}

	randIndex := rand.Intn(len(vehicles))
	rented_vehicle := vehicles[randIndex]

	wg.Add(1)
	ch <- rented_vehicle

	fmt.Printf("%s ha noleggiato il veicolo %s\n", customer.name, rented_vehicle.v_type)

	return rented_vehicle
}

func print(rented []Vehicle) {
	berline := 0
	suv := 0
	stationWagon := 0

	for _, vehicle := range rented {
		switch vehicle.v_type {
		case "Berlina":
			berline++
		case "SUV":
			suv++
		case "Station Wagon":
			stationWagon++
		}
	}

	fmt.Printf("Numero di Berline noleggiate: %d\n", berline)
	fmt.Printf("Numero di SUV noleggiati: %d\n", suv)
	fmt.Printf("Numero di Station Wagon noleggiate: %d\n", stationWagon)
}
