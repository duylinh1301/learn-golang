package main

import (
	"fmt"
	"sync"
	"time"
)

func goRoutine() {
	fmt.Println("this function run after go routine main")
}

var wg sync.WaitGroup

func main() {
	// goRoutine()
	// // time.Sleep(time.Second)
	// fmt.Printf("wait for go routine then continue run")

	// num := runtime.NumGoroutine()

	// fmt.Printf("\nnumber of go routine: %d", num)
	fmt.Println("start main")

	wg.Add(2)

	go step1()
	go step2()

	wg.Wait()

	fmt.Println("end main")

}

func step1() {
	fmt.Println("logic 1")
	time.Sleep(1)
	wg.Done()
}

func step2() {
	fmt.Println("logic 2")
	time.Sleep(2)
	wg.Done()
}
