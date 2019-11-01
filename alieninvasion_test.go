package main

import (
	"testing"
	"fmt"
)

func TestMakeMove(t *testing.T){
	var count int
	count = 0
	var cities []Cities
	cities = append(cities, Cities{"Foo", "Bar", "Qu-ux", "", "Baz", 0})
	cities = append(cities, Cities{"Bar", "", "Foo", "", "Bee", 0})
	cities = append(cities, Cities{"Qu-ux", "Foo", "", "", "Aag", 0})
	cities = append(cities, Cities{"Aag", "Baz", "", "Qu-ux", "", 0})
	cities = append(cities, Cities{"Baz", "Bee", "Aag", "Foo", "", 0})
	cities = append(cities, Cities{"Bee", "", "Baz", "Bar", "", 1})
	cities = makeMove(cities)
	for _, city := range cities {
		if city.name == "Bee" && city.alien != 0 {
			t.Errorf("Alien was not moved")
		}
		if city.name == "Bar" && city.alien != 1 {
			count = count + 1
		}
		if city.name == "Baz" && city.alien != 1 {
			count = count + 1
		}
	}
	if count == 2 {
		t.Errorf("Alien was not moved or not moved properly")
		fmt.Print(cities)
	}
}

func TestCheckCities(t *testing.T){
	var count int
	count = 0
	var cities []Cities
	cities = append(cities, Cities{"Foo", "Bar", "Qu-ux", "", "Baz", 2})
	cities = append(cities, Cities{"Bar", "", "Foo", "", "Bee", 0})
	cities = append(cities, Cities{"Qu-ux", "Foo", "", "", "Aag", 0})
	cities = append(cities, Cities{"Aag", "Baz", "", "Qu-ux", "", 0})
	cities = append(cities, Cities{"Baz", "Bee", "Aag", "Foo", "", 0})
	cities = append(cities, Cities{"Bee", "", "Baz", "Bar", "", 0})
	cities, count = checkCities(cities, count)
	for _, city := range cities {
		if city.name == "Foo" {
			t.Errorf("City was not removed")
		}
	}
	for _, city := range cities {
		if city.north == "Foo" || city.south == "Foo" || city.east == "Foo" || city.west == "Foo"{
			t.Errorf("City was not removed from possible cities")
		}
	}
}