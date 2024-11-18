package main

import "fmt"

func removeDuplicates(in, out chan string) {
	var a string
	for i := range in {
		if i != a {
			out <- i
			a = i
		}
	}
	close(out)
}

func main() {
	inputChan := make(chan string)
	outputChan := make(chan string)
	go func() {
		inputChan <- "hello"
		inputChan <- "world"
		inputChan <- "hello"
		inputChan <- "hello"
		inputChan <- "test"
		inputChan <- "alexander"
		inputChan <- "hello"
		inputChan <- "alex!"
		close(inputChan)
	}()
	go removeDuplicates(inputChan, outputChan)
	for s := range outputChan {
		fmt.Println(s)
	}
}
