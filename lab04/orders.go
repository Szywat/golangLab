package main

import (
	"math/rand"
)

func generateOrder(id int) Order {
	
	names := []string{"Martusia", "Szymon", "Michał", "Pola", "Janek"}
	prices := map[string]float64{
		"jabłko":     2.50,
		"banan":      5.0,
		"gruszka":    3.50,
		"pomarańcza": 4.0,
		"Mini Koparka Koparka Rosfix Kabina R15 Gąsienicowa 1.5T 1500KG Koop 9.5KM": 49999.99,
		"samohud": 1750.0,
		"pikachu": 4623.66,
	}
	
	randomName := names[rand.Intn(len(names))]

	keys := make([]string, 0, len(prices))
	for k := range prices {
		keys = append(keys, k)
	}

	numItems := rand.Intn(len(prices)) + 1
	used := make(map[string]bool)
	selectedItems := make([]string, 0, numItems)
	total := 0.0

	for len(selectedItems) < numItems {
		item := keys[rand.Intn(len(keys))]
		if !used[item] {
			used[item] = true
			selectedItems = append(selectedItems, item)
			total += prices[item]
		}
	}

	return Order{
		ID:           id,
		CustomerName: randomName,
		Items:        selectedItems,
		TotalAmount:  total,
	}
}
