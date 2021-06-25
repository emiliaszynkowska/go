package main

import "fmt"

func main() {
	c := push()
	for n := range pull(c) {
		fmt.Println(n)
	}
}

func push() <- chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	} ()
	return out
}

func pull(c <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	} ()
	return out
}
