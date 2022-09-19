package main

import (
	"bytes"
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	TimeUnit          = 2000 * time.Millisecond
	CookNumber        = 4
	LocalAddress      = "localhost:8087"
	DiningHallAddress = "http://localhost:8086/distribution"
)

var (
	CookRanks         = []int{3, 2, 2, 1}
	CookProficiencies = []int{4, 3, 2, 2}
	OrderList         = list.New()
)

func main() {
	go orderSender(OrderList)
	http.HandleFunc("/order", OrderHandler)
	http.ListenAndServe(LocalAddress, nil)
	//CookList := initCooks(CookRanks, CookProficiencies)

}

func orderSender(ol *list.List) {
	for {
		if ol.Len() == 0 {
			log.Println("No more orders queued up")
			time.Sleep(10 * TimeUnit)
			continue
		}

		top := ol.Front()
		var o Order = top.Value.(Order)
		b, ok := json.Marshal(o)
		o.CookingTime = time.Now().Unix()
		if ok != nil {
			log.Fatalln("Could not marshall JSON")
		}
		log.Println(o)
		ol.Remove(top)
		log.Printf("Removing an element from the list, now %v\n", ol.Len())
		if resp, ok := http.Post(DiningHallAddress, "text/json", bytes.NewBuffer(b)); ok != nil {
			fmt.Printf("Response:\t%v", resp)
			panic(ok)
		}
		log.Println()
	}
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Printf("The server only supports POST requests\n")
		return
	}
	var o Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		panic(err)
	}
	fmt.Printf("order: %+v pushed into the list\n", o)
	OrderList.PushFront(o)
	fmt.Printf("there are %v orders in the list\n", OrderList.Len())
}
