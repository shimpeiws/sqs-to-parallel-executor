package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func mainLoop(number int) {
	for {
		fmt.Printf(" mainLoop number = %d\n", number)
		cmd := exec.Command(os.Args[1], os.Args[2:]...)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		state := cmd.ProcessState
		fmt.Printf("%s\n", state.String())
		fmt.Printf(" PID %d\n", state.Pid())
		fmt.Printf(" System %v\n", state.SystemTime())
		fmt.Printf(" User %v\n", state.UserTime())
		time.Sleep(1 * time.Second)
	}
}

func main() {
	if len(os.Args) == 1 {
		return
	}
	parallelCount, err := strconv.Atoi(os.Getenv("PARALLEL_COUNT"))
	if err != nil {
		panic(err)
	}

	ch := make(chan int, parallelCount)
	for i := 0; i < parallelCount; i++ {
		go mainLoop(i)
	}

	<-ch
}
