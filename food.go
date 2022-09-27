package main

import (
	"encoding/json"
	"io/ioutil"
)

type RestaurantMenu []Dish

type Dish struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	PreparationTime  int    `json:"preparation-time"`
	Complexity       int    `json:"complexity"`
	CookingApparatus string `json:"cooking-apparatus"`
}

type KitchenDish struct {
	info      Dish
	cookedFor int
	ready     bool
}

func NewKitchenDish(id int) *KitchenDish {
	kd := new(KitchenDish)
	kd.info = CurrentMenu[id]
	kd.ready = false
	kd.cookedFor = 0
	return kd
}

func (m RestaurantMenu) ParseMenu(s string) RestaurantMenu {
	fin, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	var menu RestaurantMenu
	if err := json.Unmarshal(fin, &menu); err != nil {
		panic(err)
	}
	return menu
}
