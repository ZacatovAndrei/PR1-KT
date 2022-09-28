package main

import "time"

const (
	cReset   = "\033[0m"
	cResetNl = cReset + "\n" // simplifies one line log.printf logs
	cRed     = "\033[31m"
	cGreen   = "\033[32m"
	cYellow  = "\033[33m"
	cBlue    = "\033[34m"
	cPurple  = "\033[35m"
	cCyan    = "\033[36m"
	cGray    = "\033[37m"
	cWhite   = "\033[97m"
)

const (
	TimeUnit                     = 1000 * time.Millisecond
	CookNumber                   = 4
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
