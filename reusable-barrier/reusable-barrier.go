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


  func thread(s1 semaphore.Semaphore, s2 semaphore.Semaphore, mutex semaphore.Semaphore, count *int, id int, n int){
	for i:=0;i < 3;i++{
		
		doSomeWork(1000)

		mutex.Wait()
		(*count)++
		fmt.Println("Thread", id, "is waiting for the others")
		if (*count) == n{
			fmt.Println("Thread", id, "was the last to arrive, resume execution")
			for i := 0; i < n; i++{
				s1.Signal()
			}
		}
		mutex.Signal()

		s1.Wait()

		doSomeWork(1000)

		mutex.Wait()
		(*count)--
		fmt.Println("Thread", id, "is waiting for the others")
		if (*count) == 0{
			fmt.Println("Thread", id, "was the last to arrive, loop!")
			for i:=0;i<n;i++{
				s2.Signal()
			}
		}
		mutex.Signal()

		s2.Wait()
	}
	fmt.Println("Thread", id, "finished working")
}



func main(){
	var count int = 0
	var countp *int
	countp = &count

	var createdThreads int = 5

	var sem1 = semaphore.New(0)
	var sem2 = semaphore.New(0)
	var mutex = semaphore.New(1)
	
	for i:=0;i<createdThreads;i++{
		go thread(sem1,sem2, mutex, countp, i, createdThreads)
	}
	fmt.Scanln()
}