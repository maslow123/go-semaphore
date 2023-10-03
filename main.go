package main

import (
	"fmt"
	"time"

	"github.com/maslow123/go-semaphore/semaphore"
)

func main() {

	fmt.Println("Start")
	sem := semaphore.New(3)
	doneC := make(chan bool, 1)
	totProcess := 6
	for i := 1; i <= totProcess; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			if v == totProcess {
				doneC <- true
			}
			longRunningProcess(v)
		}(i)
	}
	<-doneC

	fmt.Println("Finish")
}
func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(2 * time.Second)
}
