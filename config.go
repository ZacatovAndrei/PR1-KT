package main

import "time"

const (
	TimeUnit                     = 2000 * time.Millisecond
	CookNumber                   = 1
	MaxFoods                     = 6
	TableNumber                  = 10
	LocalAddress                 = ":8087"
	DiningHallAddress            = "http://DiningHall:8086/distribution"
	DiningHallAddressNoContainer = "http://localhost:8086/distribution"
	MenuPath                     = "./"
	// DishesPerCookThread TODO: implement cooks working on more than one dish per proficiency thread
	// Average expectation for 1..6 dishes per order is 3.5
	// Hence at 10 tables there will be 35 orders per table on average
	// That would mean that a typical cook will have to control at least
	// math.ceil( TableNumber*(1+MaxFoods)/2 / (Sum of cook proficiencies + available cooking equipments) )
	DishesPerCookThread = 3
)
