package main

import (
	"log"
	"math"
)

type Order struct {
	OrderId        int              `json:"order_id,omitempty"`
	TableId        int              `json:"table_id,omitempty"`
	WaiterId       int              `json:"waiter_id,omitempty"`
	Items          []int            `json:"items,omitempty"`
	Priority       int              `json:"priority,omitempty"`
	MaxWait        int              `json:"max_wait,omitempty"`
	PickUpTime     int64            `json:"pick_up_time,omitempty"`
	CookingTime    int64            `json:"cooking_time,omitempty"`
	CookingDetails []map[string]int `json:"cooking_details,omitempty"`
}

func (o *Order) assignPriority() {
	//getting the order parameters
	pPriority := float64(o.Priority)
	itemLen := float64(len(o.Items))
	var avgComp, tempRank float64
	for _, item := range o.Items {
		avgComp += float64(CurrentMenu[item].Complexity)
	}
	avgComp /= itemLen
	//TODO: find a better formula for priority calculation
	tempRank = math.Round(12.0 + 6*(pPriority/5.0) - itemLen - avgComp)
	log.Printf("Current rank is %v;\n Parametres are %+v %+v %+v", tempRank, pPriority, itemLen, avgComp)
	o.Priority = int(tempRank)/2 - 1
}

func (o *Order) decompose() {
	for _, item := range o.Items {
		dishRank := CurrentMenu[item].Complexity
		switch dishRank {

		}
	}
}
