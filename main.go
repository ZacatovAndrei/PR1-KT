package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	// OrdersPending is using int32 for the sake of using the atomic package for sync
	OrdersPending = 0
	OrderList     = make([]Order, TableNumber)
	CurrentMenu   RestaurantMenu
)

// having a list of all the order items sorted by ranks

var (
	RankDish = make([][]KitchenDish, 3)
)

func main() {
	//Parsing the Menu file and initialising other things made by my sick mind
	CurrentMenu = CurrentMenu.ParseMenu(MenuPath + "menu.json")
	log.Printf("current menu :\n %+v\n", CurrentMenu)
	for i := 0; i < 3; i++ {
		RankDish[i] = make([]KitchenDish, 0, MaxOrders)
	}

	//Initializing cooks
	var CookList = make([]Cook, CookNumber)
	initCooks(CookList, CookRanks, CookProficiencies, OrderList)

	//Initializing the server side
	http.HandleFunc("/order", OrderHandler)
	if ok := http.ListenAndServe(LocalAddress, nil); ok != nil {
		panic(ok)
	}
}

func initCooks(cl []Cook, ranks []int, proficiencies []int, ol []Order) {
	for i := 0; i < CookNumber; i++ {
		go cl[i].Start(i, ranks[i], proficiencies[i], ol)
	}
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	//sanity check
	if r.Method != "POST" {
		_, _ = fmt.Fprintf(w, "The server only supports POST requests\n")
		return
	}
	//Deserializing incoming Orders
	var o Order
	if ok := json.NewDecoder(r.Body).Decode(&o); ok != nil {
		panic(ok)
	}
	log.Printf("order: %+v pushed into the list\n", o)

	//Assigning the kitchen priority based on the order parameters
	o.assignPriority()

	//Locking ListAccess mutex in the critical section
	ListAccess.Lock()
	OrderList[OrdersPending] = o
	OrdersPending++
	ListAccess.Unlock()
}

/* TODO: REMOVE the function and do something better

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
			log.FatalLn("Could not marshall JSON")
		}
		log.Println(o)
		ol.Remove(top)
		log.Printf("Removing an element from the list, now %v\n", ol.Len())
		if resp, ok := http.Post(DiningHallAddressNoContainer, "text/json", bytes.NewBuffer(b)); ok != nil {
			fmt.Printf("Response:\t%v", resp)
			panic(ok)
		}
		log.Println()
	}
}
*/
