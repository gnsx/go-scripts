package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	tasks := make(chan int)
	wg := sync.WaitGroup{}
	for worker := 0; worker < 4; worker++ {
		wg.Add(1)
		go func(tasks chan int, wg *sync.WaitGroup) {
			defer wg.Done()

			for ELement := range tasks {
				time.Sleep(time.Duration(rand.Intn(10)*200) * time.Millisecond)
				fmt.Print("\nFinished:", ELement)
			}

		}(tasks, &wg)
	}

	time.Sleep(time.Duration(1) * time.Second)
	for i := 0; i < 20; i++ {
		fmt.Print("\nSpawning i:", i)
		tasks <- i

	}

	close(tasks)
	wg.Wait()

	fmt.Print("\nBye")
}
