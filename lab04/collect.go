package main

import "fmt"

func CollectResults(results <-chan ProcessResult, failedOrders chan<- Order, done chan<- bool) {
	var (
		totalOrders int
		successOrdersCount int
		failedOrdersCount int
	)

	for result := range results {
		totalOrders++
		if result.Success {
			successOrdersCount++
			fmt.Printf("✅ Zamówienie %d zakończone powodzeniem w %s\n", result.OrderID, result.ProcessTime)
		} else {
			failedOrdersCount++
			fmt.Printf("❌ Zamówienie %d zakończone niepowodzeniem w %s\n", result.OrderID, result.ProcessTime)
			FindOrder(result.OrderID, result.CustomerName, failedOrders)
		}
	}

	successOrdersCount = successOrdersCount - failedOrdersCount
	totalOrders = totalOrders - failedOrdersCount
	
	successRate := float64(successOrdersCount) / float64(totalOrders) * 100
	failureRate := float64(failedOrdersCount) / float64(totalOrders) * 100

	fmt.Printf("\n Podsumowanie:\n")
	fmt.Printf("Wszystkich zakończomych zamówień: %d\n", totalOrders)
	fmt.Printf("Udane za pierwszym razem: %d (%.2f%%)\n", successOrdersCount, successRate)
	fmt.Printf("Nieudane za pierwszym razem: %d (%.2f%%)\n", failedOrdersCount, failureRate)

	done <- true
}

func FindOrder(orderID int, customerName string, failedOrders chan<- Order) {
	order := Order{
		ID: orderID,
		CustomerName: customerName,
	}
	failedOrders <- order
	fmt.Printf("Zamówienie %d dla %s zostało dodane do nieudanych zamówień\n", order.ID, order.CustomerName)
}