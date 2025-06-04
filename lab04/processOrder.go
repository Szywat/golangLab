package main

import ( 
	"time"
	"fmt"
	"math/rand"
	"sync"
)

func ProcessOrder(order Order) ProcessResult {
	proccesingTime := time.Duration(time.Duration(rand.Intn(2) + 1)) * time.Second
	time.Sleep(proccesingTime)

	sucess := rand.Float64() > 0.3
	var err error
	if !sucess {
		err = fmt.Errorf("error przy przetwarzaniu zamówienia %d", order.ID)
	}
	
	return ProcessResult{
		OrderID: order.ID,
		CustomerName: order.CustomerName,
		Success: sucess,
		ProcessTime: proccesingTime,
		Error: err,
	}
}

func Worker(id int, jobs <-chan Order, results chan<- ProcessResult, wg *sync.WaitGroup){
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d przetwarza zamówienie %d dla %s\n", id, job.ID, job.CustomerName)
		result := ProcessOrder(job)
		results <- result
	}
}


func RetryFailedOrders(failedOrders <-chan Order, results chan<- ProcessResult, wg *sync.WaitGroup) {
	defer wg.Done()
	
	const maxRetries = 3
	for order := range failedOrders {
		for attempt := 1; attempt <= maxRetries; attempt++ {
			fmt.Printf("Ponawianie zamówienia %d dla %s (próba %d)\n", order.ID, order.CustomerName, attempt)
			result := ProcessOrder(order)
			if result.Success {
				fmt.Printf("Zamówienie %d zakończone powodzeniem przy %d próbie\n", order.ID, attempt)
				results <- result
				break
			}	
			if attempt < maxRetries {
				time.Sleep(2 * time.Second)
			} else {
				fmt.Printf("Zamówienie %d zakończone niepowodzeniem po %d próbach.\n", order.ID, maxRetries)
			}
		}
	}
}