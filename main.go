package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/shimpeiws/sqs-to-parallel-executor/sqs"
)

func mainLoop(number int) {
	for {
		fmt.Printf(" mainLoop number = %d\n", number)
		body, err := sqs.ReceiveMessage(os.Getenv("QUEUE_URL"))
		if err != nil {
			panic(err)
		}

		var args []string
		if len(os.Args) < 3 {
			args = append(os.Args, body)
		} else {
			position := 2
			args = append(os.Args[:position+1], os.Args[position:]...)
			args[position] = body
		}
		cmd := exec.Command(args[1], args[2:]...)
		res, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Printf("response = %s\n", res)
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
	defer close(ch)
	for i := 0; i < parallelCount; i++ {
		go mainLoop(i)
	}

	<-ch
}
