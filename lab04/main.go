package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)

func main() {
	orders := make(chan Order, 100)
	results := make(chan ProcessResult, 100)
	failedOrders := make(chan Order, 100)
	done := make(chan bool)

	go CollectResults(results, failedOrders, done)

	var wg sync.WaitGroup
	numberOfWorkers := 5
	for i:= 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go Worker(i, orders, results, &wg)
	}

	go func() {
		for i := 1; i <= 5; i++ {
			order := generateOrder(i)
			fmt.Printf("Wygenerowano zamówienie %d dla %s za %.2f zł\n", order.ID, order.CustomerName, order.TotalAmount)
			orders <- order
			
			time.Sleep(time.Duration(rand.Intn(2) + 1) * time.Second)
		}
		
		close(orders)
	}()

	wg.Wait()
	
	println("\nWszystkie zamówienia zostały przetworzone")
	println("Podjęcie próby nad przetworzeniem nieudanych zamówień")
	
	wg.Add(1)
	go RetryFailedOrders(failedOrders, results, &wg)
	
	close(failedOrders)
	wg.Wait()

	close(results)
	<-done

	fmt.Println("Wszystkie gorutyny zakończyły swoją pracę")
}