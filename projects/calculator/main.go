package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ans := make(chan int)
	go func(a chan int) {
		defer close(a)
		select {
		case b := <-firstChan:
			a <- b * b
		case d := <-secondChan:
			a <- d * 3
		case <-stopChan:
			return
		}
	}(ans)
	return ans
}

func main() {
	first, second := make(chan int), make(chan int)
	stopChan := make(chan struct{})
	calc := calculator(first, second, stopChan)
	first <- 4
	//second <- 2
	close(stopChan)
	fmt.Println(<-calc)
}
