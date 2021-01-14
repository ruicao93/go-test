package main

import (
	"fmt"
	"sync"
)

func Test1() {
	lock := sync.Mutex{}
	cond := sync.NewCond(&lock)
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		for count == 0 {
			cond.Wait()
		}
		fmt.Print("baba")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		count = 1
		fmt.Print("ali")
		wg.Done()
	}()
	wg.Wait()
}

func Test2() {
	ch := make(chan int)
	stop := make(chan int)
	go func() {
		<-ch
		fmt.Print("baba")
	}()
	go func() {
		fmt.Print("ali")
		ch <- 1
		stop <- 1
	}()
	<-stop
}

func main() {
	Test1()
	fmt.Println()
	Test2()
}
