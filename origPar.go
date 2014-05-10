package main

import "fmt"
import "sync"

func Par(processes ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(processes))
	for _, v := range processes {
		go func(v func()) {
			defer wg.Done()
			v()
		}(v)
	}
	wg.Wait()
}

func main() {
	c := make(chan int)
	Par(
		func() {
			c <- 100
		},
		func() {
			v := <-c
			fmt.Println(v) 
		}	)
}
