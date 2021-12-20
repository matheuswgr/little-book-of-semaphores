package main

import (
	"fmt"
	"time"
	"math/rand"
	"little-book-of-semaphores/synchronization-primitives"
  )

  func threadA(s semaphore.Semaphore){
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
		
		fmt.Println("Hey, thread A here")
		s.Signal()
  }

  func threadB(s semaphore.Semaphore){
		s.Wait()
		amt := time.Duration(rand.Intn(10))
		time.Sleep(time.Millisecond * amt)
		fmt.Println("Hey, thread B here")
  }
  
  func main() {
	var s = semaphore.New(0)
	go threadB(s)
	go threadA(s)
	var input string
	fmt.Scanln(&input)
  }