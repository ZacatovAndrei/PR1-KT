//TODO: make an OrderScheduler class that would give orders to cooks

package main

import (
	"log"
	"sync"
	"time"
)

//variables specific to Cooks
var (
	ListAccess        sync.Mutex
	CookRanks         = []int{3, 2, 2, 1}
	CookProficiencies = []int{4, 3, 2, 2}
)

type Cook struct {
	id, rank, proficiency int
	currentOrder          *Order
	testCh                chan bool
}

func (c *Cook) Start(id, rank, proficiency int, orderList []Order) {
	c.Init(id, rank, proficiency)
}

func (c *Cook) Init(i, r, p int) {
	c.id = i
	c.rank = r
	c.proficiency = p
	c.currentOrder = nil
	log.Printf("Initialised Cook #%v", c.id)
}

func (c *Cook) takeOrder(ol []Order) {
}

func (c *Cook) cookItem(id int) {
	time.Sleep(time.Duration(CurrentMenu[id].PreparationTime) * TimeUnit)
}
