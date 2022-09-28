package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	OrdersPending = 0
	OrderList     = make([]Order, TableNumber)
	CurrentMenu   RestaurantMenu
)

func main() {
	//Parsing the Menu file
	CurrentMenu = CurrentMenu.ParseMenu(MenuPath + "menu.json")
	log.Printf("current menu :\n %+v\n\n", CurrentMenu)

	//Initializing cooks
	var CookList = make([]Cook, CookNumber+1)
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
	log.Printf(cGreen+"order: %+v pushed into the list"+cResetNl, o)
	//Assigning the kitchen priority based on the order parameters
	o.assignPriority()
	//Locking ListAccess mutex in the critical section
	ListAccess.Lock()
	OrderList[OrdersPending] = o
	OrdersPending += 1
	ListAccess.Unlock()
}
