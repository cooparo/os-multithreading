package main

import "sync"

func countChar(s rune, c rune, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	if s == c {
		temp := <- ch
		temp++
		ch <- temp
	}
}

func main() {

	const str = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
	const c = 'c'

	ch := make(chan int, 1)
	ch <- 0

	wg := sync.WaitGroup{}

	// Start a thread for each char in str
	for i := 0; i < len(str); i++ {
		wg.Add(1)
		go countChar(rune(str[i]), c, ch, &wg)
	}

	wg.Wait()

	count := <- ch
	close(ch)

	println("Count of char", string(c), "is", count)

}
