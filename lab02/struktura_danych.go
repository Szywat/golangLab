package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Location struct {
	latitude   float64
	longitutde float64
}

type Data struct {
	code            int
	province_state  string
	admin_location  string
	date            string
	total_death     int
	total_confirmed int
	location        Location
}

func readCSVLineByLine(filepath string) []Data {
	file, err := os.Open(filepath) // "dane.csv"

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var data []Data

	for i, r := range records {
		if i == 0 { // pominięcie nagłówka
			continue
		}

		code, _ := strconv.Atoi(r[0])
		total_death, _ := strconv.Atoi(r[4])
		total_confirmed, _ := strconv.Atoi(r[5])

		var lat, long float64
		locationParts := strings.Split(r[6], ", ")
		if len(locationParts) == 2 {
			lat, _ = strconv.ParseFloat(locationParts[0], 64)
			long, _ = strconv.ParseFloat(locationParts[1], 64)
		}

		data = append(data, Data{code, r[1], r[2], r[3], total_death, total_confirmed, Location{lat, long}})
	}
	return data
}

func sorting(data []Data) {
	sort.Slice(data, func(i, j int) bool { return data[i].total_death > data[j].total_death })
	fmt.Println("Total death:")
	for _, d := range data[:5] {
		fmt.Println(d)
	}
	fmt.Printf("\n")

	sort.Slice(data, func(i, j int) bool { return data[i].location.latitude > data[j].location.latitude })
	fmt.Println("Longitude:")
	for _, d := range data[:5] {
		fmt.Println(d)
	}
}

func statistic(data []Data) {
	casesByDate := make(map[string]int)
	for _, d := range data {
		casesByDate[d.date] += d.total_death
	}

	maxCases := 0
	maxDate := ""

	for date, cases := range casesByDate {
		if cases > maxCases {
			maxCases = cases
			maxDate = date
		}
	}
	fmt.Printf("\nDzień z największą liczbą śmierci: %s (%d)\n", maxDate, maxCases)

}

func main() {
	new_data := readCSVLineByLine("dane.csv")

	sorting(new_data)
	statistic(new_data)
}
