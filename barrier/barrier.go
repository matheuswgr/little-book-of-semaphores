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

func thread(s semaphore.Semaphore,mutex semaphore.Semaphore, count *int, id int, n int){
	
	doSomeWork(250)
	
	mutex.Wait()
	(*count)++
	fmt.Println("Thread", id, "is waiting to access the critical section")
	if (*count) == (n){
		fmt.Println("Thread", id, "releasing execution")
		s.Signal()
	}
	mutex.Signal()

	s.Wait()
	s.Signal()

	doSomeWork(250)	
	fmt.Println("Thread", id, "has accessed the critical section")
}

func main(){
	var count int = 0
	var countp *int
	countp = &count

	var createdThreads int = 10

	var sem = semaphore.New(0)
	var mutex = semaphore.New(1)
	
	for i:=0;i<createdThreads;i++{
		go thread(sem, mutex, countp, i, createdThreads)
	}
	fmt.Scanln()
}