package semaphore

type Semaphore struct{
	sem chan struct{}
}

func New(initialValue int) Semaphore{
	newSemaphore := Semaphore{
		sem: make(chan struct{}, 10000),
	}

	for i:=0;i<initialValue;i++{
		newSemaphore.Signal()
	}

	return newSemaphore
}

func (s Semaphore) Signal(){
	s.sem <- struct{}{}
}

func (s Semaphore) Wait(){
	<- s.sem
}