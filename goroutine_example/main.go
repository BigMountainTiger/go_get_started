package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	f1 := func() {
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

		fmt.Println()
	}

	f1()

	f2 := func() {
		var wg sync.WaitGroup
		var lock sync.Mutex

		// Different Mutex instance do not block each other
		// var lock_1 sync.Mutex

		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Second)

			fmt.Println("Trying to aquire lock in No.1")
			lock.Lock()
			fmt.Println("Aquired lock in No.1")

			defer func() {
				lock.Unlock()
				fmt.Println("Released lock in No.1")
			}()
			defer wg.Done()

			fmt.Println("Finished No.1")
		}()

		wg.Add(1)
		go func() {

			fmt.Println("Trying to aquire lock in No.2")
			lock.Lock()
			fmt.Println("Aquired lock in No.2")

			defer func() {
				lock.Unlock()
				fmt.Println("Released lock in No.2")
			}()
			defer wg.Done()

			time.Sleep(3 * time.Second)
			fmt.Println("Finished No.2")
		}()

		wg.Wait()
		fmt.Println("Finished main func")

		fmt.Println("")
		fmt.Println("It has been tested => Different Mutex instance do not block each other")

		fmt.Println("Mutex and WaitGroup are commonly used in thread management")
	}

	f2()

	func() {
		fmt.Println()
		fmt.Println("Locks are not pointers, so it is better to use lock pointer")
		fmt.Println("so we know we which exact lock we are aquiring")

		var wg sync.WaitGroup
		var lock sync.Mutex

		lk1 := &lock
		lk2 := &lock

		wg.Add(2)
		fmt.Printf("The lock pointers are the same => %v\n", lk1 == lk2)

		go func() {
			lk1.Lock()
			defer wg.Done()
			defer lk1.Unlock()

			fmt.Println("lock1 is aquired")
			time.Sleep(3 * time.Second)

		}()

		go func() {
			lk2.Lock()
			defer wg.Done()
			defer lk2.Unlock()

			fmt.Println("lock2 is aquired")
			time.Sleep(3 * time.Second)

		}()

		wg.Wait()
		fmt.Println("Finished both go-routines")

	}()

	func() {
		fmt.Println()
		fmt.Println("Testing if initializing sync.Mutex is needed in a struct")

		type T struct {
			text_1 string
			text_2 string
			lock   sync.Mutex
		}

		t := T{
			text_1: "Text 1",
			text_2: "Text 2",
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			t.lock.Lock()
			defer t.lock.Unlock()

			fmt.Printf("lock aquired, printing - %s\n", t.text_1)
			time.Sleep(3 * time.Second)

		}()

		go func() {
			defer wg.Done()

			t.lock.Lock()
			defer t.lock.Unlock()

			fmt.Printf("lock aquired, printing - %s\n", t.text_2)
			time.Sleep(3 * time.Second)

		}()

		wg.Wait()
		fmt.Println("There is no need to initialize sync.Mutex in a struct")

	}()

	func() {

		fmt.Println()
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Printf("Slept for 3 seconds. This will not get printed\n")
		}()

		fmt.Println("Go routines are daemon by themselves")
	}()

}
