package main

import(
	"fmt"
	"time"
	"math/rand"
	"little-book-of-semaphores/synchronization-primitives"
)

var followers int = 0
var leaders int = 0

func doSomeWork(maximumLatency int){
	amt := time.Duration(rand.Intn(maximumLatency))
	time.Sleep(time.Millisecond * amt)
  }

func leaderArrival(leadersQueue semaphore.Semaphore, followersQueue semaphore.Semaphore, mutex semaphore.Semaphore, id int, rendezvous semaphore.Semaphore){
	doSomeWork(500)
	mutex.Wait()
	fmt.Println("Leader",id, "has arrived!")
	if followers > 0{
		followers--
		followersQueue.Signal()
	}else{
		leaders++
		mutex.Signal()
		fmt.Println("Leader", id, "is waiting for a follower")
		leadersQueue.Wait()
	}
	fmt.Println("Leader", id, "has a follower")
	rendezvous.Wait()
	mutex.Signal()
	doSomeWork(500)
}

func followerArrival(leadersQueue semaphore.Semaphore, followersQueue semaphore.Semaphore, mutex semaphore.Semaphore, id int, rendezvous semaphore.Semaphore){
	doSomeWork(500)
	mutex.Wait()
	fmt.Println("Follower",id, "has arrived!")
	if leaders > 0{
		leaders--
		leadersQueue.Signal()
	}else{
		followers++
		fmt.Println("Follower", id, "is waiting for a leader")
		mutex.Signal()
		followersQueue.Wait()
	}
	fmt.Println("Follower", id, "has a leader")
	rendezvous.Signal()
	doSomeWork(500)
}

func main(){
	var followersQueue = semaphore.New(0)
	var leadersQueue = semaphore.New(0)
	var rendezvous = semaphore.New(0)
	var mutex = semaphore.New(1)

	createdThreads := 4
	
	for i:=0;i<createdThreads;i++{
		go leaderArrival(leadersQueue,followersQueue, mutex, i, rendezvous)
		go followerArrival(leadersQueue,followersQueue, mutex, i, rendezvous)
	}
	fmt.Scanln()
}