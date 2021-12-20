package main

import(
	"fmt"
	"time"
	"math/rand"
	"little-book-of-semaphores/synchronization-primitives"
)

func doSomeWork(maximumLatency int){
	amt := time.Duration(rand.Intn(maximumLatency))
	time.Sleep(time.Millisecond * amt)
  }

func thread(s semaphore.Semaphore, count *int, id int){
	s.Wait()
	fmt.Println("Thread", id,"is running")
	doSomeWork(250)
	(*count)++
	fmt.Println("Thread", id, "says count =", *count)
	s.Signal()
}

func main(){
	var count int = 0
	var countp *int;
	countp = &count

	var maximumConcurrentThreads = 10
	var createdThreads = 15

	var sem = semaphore.New(maximumConcurrentThreads)

	fmt.Println("Expected value for count:", createdThreads)
	
	for i:=0;i<createdThreads;i++{
		go thread(sem, countp, i)
	}
	fmt.Scanln()
}