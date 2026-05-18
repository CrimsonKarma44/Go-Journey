package main

import (
	"fmt"
	"time"
)

func f(sample2 chan int, sample1 chan string) {
	var i int
	for i = 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		sample2 <- i
	}
	close(sample2)

	sample1 <- "Hello World"
	close(sample1)
}

func c(text chan string) {
	fmt.Println(<-text)
}

func main() {
	var sampleChannel = make(chan string)
	var sampleChannel2 = make(chan int)

	go f(sampleChannel2, sampleChannel)
	go c(sampleChannel)
	for {
		i, ok := <-sampleChannel2
		if !ok {
			break
		}
		fmt.Println(i, ok)
	}
}
