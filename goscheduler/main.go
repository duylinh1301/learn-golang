package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(4) // add go routine

	runtime.GOMAXPROCS(0)

	go goroutine1()
	go goroutine2()
	go goroutine3()
	go goroutine4()

	wg.Wait()
	fmt.Print("end main")

}

func goroutine1() {
	fmt.Println("goroutine 1")
	wg.Done()
}

func goroutine2() {
	time.Sleep(2 * time.Second)
	fmt.Println("goroutine 2")
	wg.Done()
}

func goroutine3() {
	fmt.Println("goroutine 3")
	wg.Done()
}

func goroutine4() {
	time.Sleep(3 * time.Second)
	fmt.Println("goroutine 4")
	wg.Done()
}
