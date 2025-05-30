package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {

			for job := range jobChan {
				r_num := job.RandNum

				sum := 0
				for r_num != 0 {
					sum += r_num % 10
					r_num /= 10
				}

				r := &Result{
					job: job,
					sum: sum,
				}

				resultChan <- r
			}

		}(jobChan, resultChan)
	}
}

func main() {
	job_chan := make(chan *Job, 100)
	result_chan := make(chan *Result, 100)

	createPool(64, job_chan, result_chan)

	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("Job %d, value %d, has sum %d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(result_chan)

	id := 0

	for {
		id++

		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}

		job_chan <- job

		if id == 100000 {
			break
		}
	}
}
