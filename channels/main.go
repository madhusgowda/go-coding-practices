package main

func main() {
	ch := make(chan struct{})
	println("Channel initialized")
	//define workerRoutine
	go workerRoutine(ch)

	// Send signal to worker goroutine
	ch <- struct{}{}

	// Receive a message from the workerRoutine.
	<-ch
	println("Main File Signal Received")
}

func workerRoutine(ch chan struct{}) {
	// Receive message from main program.
	<-ch
	println("Worker Routine Signal Received")

	// Send a message to the main program.
	close(ch)
}
