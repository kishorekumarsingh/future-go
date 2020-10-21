package main

import (
	"fmt"
	"getmega/future"
	"time"
)

// CancelDemo demonstrates the cancel use case of future package
// by interrupting the task after starting.
func CancelDemo() {
	fmt.Println("CancelDemo started.")
	task := future.PrintInt()
	fmt.Println("Running : ", task.Running())
	task.Cancel()
	fmt.Println("Running : ", task.Running())
	fmt.Println("CancelDemo finished.\n\n")
}

// ResultDemo demonstrates the result use case of future package
// by letting the task execute till completion.
func ResultDemo() {
	fmt.Println("ResultDemo started.")
	task := future.PrintInt()
	fmt.Println("Running : ", task.Running())
	time.Sleep(12 * time.Second)
	fmt.Println("Running : ", task.Running())
	fmt.Println("ResultDemo finished.\n\n")
}

// RunningDemo demonstrates the running use case of future package
// by letting the task execute till completion and checking the status of task at various points.
func RunningDemo() {
	fmt.Println("RunningDemo started.")
	task := future.PrintInt()
	time.Sleep(2 * time.Second)
	fmt.Println("Running : ", task.Running())
	time.Sleep(3 * time.Second)
	fmt.Println("Running : ", task.Running())
	task.Cancel()
	fmt.Println("Running : ", task.Running())
	fmt.Println("RunningDemo finished.\n\n")
}

func main() {
	CancelDemo()
	ResultDemo()
	RunningDemo()
}
