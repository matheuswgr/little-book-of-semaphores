package main

import (
	"fmt"
	"time"
	"math/rand"
	"little-book-of-semaphores/synchronization-primitives"
  )

  func doSomeWork(maximumLatency int){
	amt := time.Duration(rand.Intn(maximumLatency))
	time.Sleep(time.Millisecond * amt)
  }

  func threadA(s1 semaphore.Semaphore, s2 semaphore.Semaphore){		
		go doSomeWork(500)
		fmt.Println("Hey, thread A did work 1 here")
		s1.Signal()

		s2.Wait()
		go doSomeWork(500)
		fmt.Println("Hey, thread A did work 2 here")
  }

  func threadB(s1 semaphore.Semaphore, s2 semaphore.Semaphore){
		go doSomeWork(500)
		fmt.Println("Hey, thread B did work 1 here")
		s2.Signal()

		s1.Wait()
		go doSomeWork(500)
		fmt.Println("Hey, thread B did work 2 here")
  }
  
  func main() {
	var s1 = semaphore.New(0)
	var s2 = semaphore.New(0)

	go threadB(s1, s2)
	go threadA(s1, s2)	

	var input string
	fmt.Scanln(&input)
  }