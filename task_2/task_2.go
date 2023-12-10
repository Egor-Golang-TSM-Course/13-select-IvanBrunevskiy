package task_2

import (
	"fmt"
	"math/rand"
	"time"
)

type Restaurant struct {
	RestaurantName string
	CountSoldDish  int
}

func StartSale() {
	restaurants := make(map[string]*Restaurant)
	restaurants["rest1"] = &Restaurant{RestaurantName: "tempo"}
	restaurants["rest2"] = &Restaurant{RestaurantName: "dodo"}
	restaurants["rest3"] = &Restaurant{RestaurantName: "dominos"}

	results := make(chan Restaurant)
	go CentralOffice(results)
	for _, value := range restaurants {

		go func(rest *Restaurant) {
			for {
				randomNumber := GenerateCountSoldDish()
				time.Sleep(2 * time.Second)
				rest.CountSoldDish += randomNumber
				results <- *rest
			}
		}(value)

	}
	time.Sleep(8 * time.Second)
	close(results)
}

func CentralOffice(results chan Restaurant) {
	for result := range results {
		fmt.Printf("Продажи ресторана: %+v\n", result)
	}
}

func GenerateCountSoldDish() int {
	return rand.Intn(101)
}
