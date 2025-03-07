package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// GeneratePESEL: geneuje numer PESEL
// Parametry:
// - birthDate: time.Time: reprezentacja daty urodzenia
// - płeć: znak "M" lub "K"
// Wyjscie:
//Tablica z cyframi numeru PESEL

func GenerujPESEL(birthhdate time.Time, gender string) [11]int {
	var cyfryPESEL [11]int

	year := birthhdate.Year()
	month := int(birthhdate.Month())
	day := birthhdate.Day()

	cyfryPESEL[0] = year % 100 / 10
	cyfryPESEL[1] = year % 10

	if (year <= 1899) && (year >= 1800) {
		month += 80
	} else if (year <= 2099) && (year >= 2000) {
		month += 20
	} else if (year <= 2199) && (year >= 2100) {
		month += 40
	} else if (year <= 2299) && (year >= 2200) {
		month += 60
	}
	
	cyfryPESEL[2] = month / 10
	cyfryPESEL[3] = month % 10

	cyfryPESEL[4] = day / 10
	cyfryPESEL[5] = day % 10

	randomSerial := rand.IntN(900) + 100

	cyfryPESEL[6] = randomSerial / 100
	cyfryPESEL[7] = (randomSerial % 100) / 10
	cyfryPESEL[8] = randomSerial % 10

	randGenderNum := rand.IntN(5) * 2

	switch gender {
	case "M", "m":
		cyfryPESEL[9] = randGenderNum + 1
	case "K", "k":
		cyfryPESEL[9] = randGenderNum
	default:
		cyfryPESEL[9] = -1
	}

	table := [4]int{1,3,7,9}
	for i := 0; i < len(cyfryPESEL) -1; i++ {

		cyfryPESEL[10] += (cyfryPESEL[i] * table[i % 4]) % 10

	}

	cyfryPESEL[10] = (10 - (cyfryPESEL[10] % 10)) % 10
	
	return cyfryPESEL
}

func WeryfikujPESEL(cyfryPESEL [11]int) bool {

	var czyPESEL bool
	czyPESEL = true
	// PESEL: Czy wartości większe od 10 LUB mniejsze od 0
	for i := 0; i < len(cyfryPESEL); i++ {
		if (cyfryPESEL[i] < 0) || (cyfryPESEL[i] > 9) {
			czyPESEL = false
			return czyPESEL
		} 
	}

	// CYFRA KONTROLNA: Czy się zgadza
	table := [4]int{1,3,7,9}
	val := 0
	for i := 0; i < len(cyfryPESEL) -1; i++ {
		
		val += (cyfryPESEL[i] * table[i % 4]) % 10

	}

	val = (10 - (val % 10)) % 10
	
	if val != cyfryPESEL[10] {
		czyPESEL = false
		return czyPESEL
	}

	return czyPESEL
}

func main() {
	birthDate := time.Date(2003, 15, 29, 0, 0, 0, 0, time.UTC)
	pesel := GenerujPESEL(birthDate, "M")
	fmt.Println("Wygenerowany PESEL:", pesel)

	fmt.Println("Czy numer PESEL jest poprawny:", WeryfikujPESEL(pesel))

}