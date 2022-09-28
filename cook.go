//TODO: make an OrderScheduler class that would give orders to cooks

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
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
	for {
		if err := c.takeOrder(orderList); err != nil {
			time.Sleep(3 * TimeUnit)
			continue
		}
		c.cookOrder()
		c.sendOrderToDH()
	}
}

func (c *Cook) Init(i, r, p int) {
	c.id = i
	c.rank = r
	c.proficiency = p
	c.currentOrder = nil
	log.Printf(cGreen+"Initialised Cook #%v"+cResetNl, c.id)
}

func (c *Cook) takeOrder(ol []Order) error {
	ListAccess.Lock()
	defer ListAccess.Unlock()
	if OrdersPending < 1 {
		return errors.New("no orders in the kitchen")
	}
	c.currentOrder = &ol[OrdersPending-1]
	OrdersPending--
	fmt.Printf("Orders pending:%v\nOrder %v picked by Cook %v\n", OrdersPending, c.currentOrder.OrderId, c.id)
	return nil
}

func (c *Cook) cookOrder() {
	if c.currentOrder == nil {
		time.Sleep(TimeUnit)
		return
	}
	time.Sleep(time.Duration(c.currentOrder.MaxWait + rand.Intn(15)))
	c.currentOrder.CookingTime = time.Now().Unix()
	return
}
func (c *Cook) sendOrderToDH() {
	b, ok := json.Marshal(c.currentOrder)
	if ok != nil {
		log.Fatalln("Could not marshall JSON")
	}
	if resp, ok := http.Post(DiningHallAddress, "text/json", bytes.NewBuffer(b)); ok != nil {
		fmt.Printf("Response:\t%v", resp)
		panic(ok)
	}
	log.Printf(cCyan+"Order %v sent to the Dining Hall successfully"+cResetNl, c.currentOrder.OrderId)
	c.currentOrder = nil
}
