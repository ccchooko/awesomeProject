//workerpool.go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//工作线程
func workerPool(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("start job", j)
		time.Sleep(time.Second) //模拟耗时工作
		fmt.Println("finish job", j)
		results <- j
	}
}
func init() {
	wg = sync.WaitGroup{}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)    //
	results := make(chan int, numJobs) //

	go workerPool(jobs, results)

	for i := 1; i <= numJobs; i++ {
		wg.Add(1)
		jobs <- i
	}

	go func() {
		for r := range results {
			fmt.Println("results: ", r)
			wg.Done()
		}
	}()

	wg.Wait()
	//for i := 1; i <= numJobs; i++ {
	//	jobs <- i
	//}
	//close(jobs)
	////for r := range results {
	////	fmt.Println("results :", r)
	////}
	//time.Sleep(20 * time.Second)
}
