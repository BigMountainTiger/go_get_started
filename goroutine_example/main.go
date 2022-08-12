package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
			fmt.Println("Recovered")
		}
	}()

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ip := new(int)
	*ip = 3

	fmt.Println(*ip)

}
