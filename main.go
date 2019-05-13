package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/shimpeiws/sqs-to-parallel-executor/sqs"
)

func pollingInterval() time.Duration {
	defaultInterval := 1 * time.Second
	if value, ok := os.LookupEnv("POLLING_INTERVAL_SECONDS"); ok {
		var intervalSecond int64
		intervalSecond, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return defaultInterval
		}
		return time.Duration(intervalSecond) * time.Second
	}
	return defaultInterval
}

func mainLoop(number int) {
	for {
		log.Printf(" mainLoop number = %d\n", number)
		body, err := sqs.ReceiveMessage(os.Getenv("QUEUE_URL"))
		if err != nil {
			log.Panic(err)
		}

		if len(body) == 0 {
			time.Sleep(pollingInterval())
			continue
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
			log.Panic(err)
		}
		log.Printf("response = %s\n", res)
		time.Sleep(pollingInterval())
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
