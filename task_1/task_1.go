package task_1

import (
	"fmt"
	"time"
)

type Chef struct {
	Name     string
	CookTime int
}

const RestorauntWorkTime = 10

func StartRestaurant() {
	chefs := make(map[string]Chef)

	chefs["chef1"] = Chef{Name: "Ivan", CookTime: 3}
	chefs["chef2"] = Chef{Name: "Svetlana", CookTime: 4}

	for _, value := range chefs {
		go func(chef Chef) {
			for {
				fmt.Printf("Chef %s, started cook \n", chef.Name)
				time.Sleep(time.Duration(chef.CookTime) * time.Second)
				fmt.Printf("Chef %s, finished cook \n", chef.Name)
			}
		}(value)
	}

	time.Sleep(RestorauntWorkTime * time.Second)
}
