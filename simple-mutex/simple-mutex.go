package main

import(
	"fmt"
	"time"
	"math/rand"
	"little-book-of-semaphores/synchronization-primitives"
)

func threadA(s semaphore.Semaphore, count *int){
	s.Wait()
	time.Sleep(time.Millisecond*time.Duration(rand.Intn(250)))
	(*count)++
	fmt.Println("Thread A says count =", *count)
	s.Signal()
}

func threadB(s semaphore.Semaphore, count *int){
	s.Wait()
	time.Sleep(time.Millisecond*time.Duration(rand.Intn(250)))
	(*count)++
	fmt.Println("Thread B says count =", *count)
	s.Signal()
}

func main(){
	var count int = 0
	var countp *int;
	countp = &count
	var sem = semaphore.New(1)
	fmt.Println("Expected value for count: 20")
	for i:=0;i<10;i++{
		go threadB(sem, countp)
		go threadA(sem, countp)
	}
	fmt.Scanln()
}