package main

import (
	"sync"
	"time"
	"fmt"
)

func sleepFun(sec time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(sec * time.Second)
	fmt.Println("goroutine exit")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sleepFun(1, &wg)
	wg.Wait()
	fmt.Println("Main goroutine exit")

}