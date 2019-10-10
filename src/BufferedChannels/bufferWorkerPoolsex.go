package BufferedChannels

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

type Job struct {
	id int
	randomno int
}
type Result struct {
	job Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)  
var results = make(chan Result, 10)
// this function returns sum of psudo numbers in a given number =>
// 123 => (1 + 2 + 3) = 6
func digits(number int) int {
	sum := 0
	no := number
	for no!= 0{
		digit := no % 10
		sum = sum + digit 
		no = no / 10
	}
	time.Sleep(2 * time.Second)
	return sum
}


func worker(wg *sync.WaitGroup)  {
	for job:= range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int)  {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int)  {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job:= Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool)  {
	for result:= range results {
		fmt.Printf("Job id %d, input random number %d, sum of digits %d", result.job.id, result.job.randomno, result.sumOfDigits)
	}
	done<- true
}

func Mainn2()  {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<- done
	endTime:= time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}