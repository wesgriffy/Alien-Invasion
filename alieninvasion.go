package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"strings"
	"os"
	"strconv"
)

type Cities struct{
	name string
	north string
	south string
	east string
	west string
	alien int
}

func after(value string, a string) string {
    // Handles edge case where city is last listed, gets substring after string
    pos := strings.LastIndex(value, a)
    if pos == -1 {
        return ""
    }
    adjustedPos := pos + len(a)
    if adjustedPos >= len(value) {
        return ""
    }
    return value[adjustedPos:len(value)]
}

func GetStringInBetween(str string, start string, end string) (result string) {
	//Gets string in between two substrings, assigns values to struct
	//used to assign cities north, south, east and west
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : s+e]
}

func main(){
	var cities []Cities
	var count int
	count = 0
	var lines []string
	//recieves number of aliens from command line
	nAliens, er := strconv.ParseInt(os.Args[1],0, 0)
	check(er)
	//gets file from command line
	file := os.Args[2]
	dat, err := os.Open(file)
	check(err)
	scanner := bufio.NewScanner(dat)
	//reads lines from file, gets list of city names
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		l := strings.Split(scanner.Text(), " ")
		cities = append(cities, Cities{l[0], "","","","",0})
	}
	cities = makeCities(cities, lines, int(nAliens))
	//lets simulation run so each alien makes 10,000 moves
	for a := 0; a<10000; a++{
		cities = makeMove(cities)
		cities, count = checkCities(cities, count)
	}
	reconstruct(cities)
}

func makeCities(cities []Cities, lines []string, nAliens int) ([]Cities){
	//Assigns cities to struct
	for i:= 0; i<len(lines); i++{
		if(strings.Contains(lines[i], "north=")){
			cities[i].north = GetStringInBetween(lines[i], "north=", " ")
			if(cities[i].north == ""){
				cities[i].north = after(lines[i], "north=")
			}
		}else{
			cities[i].north = ""
		}
		if(strings.Contains(lines[i], "south=")){
			cities[i].south = GetStringInBetween(lines[i], "south=", " ")
			if(cities[i].south == ""){
				cities[i].south = after(lines[i], "south=")
			}
		}else{
			cities[i].south = ""
		}
		if(strings.Contains(lines[i], "east=")){
			cities[i].east = GetStringInBetween(lines[i], "east=", " ")
			if(cities[i].east == ""){
				cities[i].east = after(lines[i], "east=")
			}
		}else{
			cities[i].east = ""
		}
		if(strings.Contains(lines[i], "west=")){
			cities[i].west = GetStringInBetween(lines[i], "west=", " ")
			if(cities[i].west == ""){
				cities[i].west = after(lines[i], "west=")
			}
		}else{
			cities[i].west = ""
		}
	}
	for j := 0; j<nAliens; j++{
		//Places an alien in a different city, assuming the number of aliens is less than the number of cities
		cities[j].alien++
	}
	return cities
}
//makes a move for each alien in each city, returns cities with aliens moved
func makeMove(cities []Cities) ([]Cities){
	l := [4]string{"north", "south", "east", "west"}
	//l will be the 4 random options for the alien to move
	for i := 0; i<len(cities); i++{
		if(cities[i].alien == 0){
			continue
		}
		rand.Seed(time.Now().UnixNano())
		direction := fmt.Sprint(l[rand.Intn(len(l))])
		switch direction {
			case "north" :
				if(cities[i].north != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].north {
							cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].south != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].south {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].east != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].east {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].west != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].west {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}
			case "south" :
				if(cities[i].south != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].south {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].east != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].east {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].west != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].west {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].north != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].north {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}
			case "east":
				if(cities[i].east != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].east {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].west != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].west {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].north != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].north {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].south != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].south {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}
			case "west":
				if(cities[i].west != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].west {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].north != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].north {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].south != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].south {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}else if(cities[i].east != ""){
					for j := 0; j<len(cities); j++ {
    					if cities[j].name == cities[i].east {
        					cities[j].alien = cities[j].alien + 1
    					}
					}
					cities[i].alien--
				}
		}
	}
	return cities
}
//Checks the cities to see if they have multiple aliens, if they do, they fight
func checkCities(cities []Cities, count int) ([]Cities, int){
	var v [] string
	for i := 0; i<len(cities); i++{
		if(cities[i].alien == 2){
			fmt.Println(fmt.Sprint(cities[i].name, " has been destroyed by alien", count, " and alien",count+1))
			count = count + 2
			v = append(v, cities[i].name)
		}
		if(cities[i].alien > 2){
			fmt.Print(fmt.Sprint(cities[i].name, " has been destroyed by"))
			for b := 0; b<cities[i].alien-1; b++{
				fmt.Print(fmt.Sprint(" alien", count, ","))
				count = count + 1
			}
			fmt.Println(fmt.Sprint(" and alien", count))
			count = count + 1
			v = append(v, cities[i].name)
		}
	}
	for j := 0; j<len(v); j++{
		//removes cities that were destroyed
		cities = remove(cities, v[j])
		//removes cities from list of possible moves aliens can make
		cities = remover(cities, v[j])
	}
	return cities, count
}

func remover(cities []Cities, s string) ([]Cities){
	//removes cities from list of possible moves aliens can make
	for i:=0; i<len(cities); i++{
		if(cities[i].north == s){
			cities[i].north = ""
		}
		if(cities[i].south == s){
			cities[i].south = ""
		}
		if(cities[i].east == s){
			cities[i].east = ""
		}
		if(cities[i].west == s){
			cities[i].west = ""
		}
	}
	return cities
}

func remove(cities []Cities, s string) ([]Cities){
	//removes cities that were destroyed
	for i:=0; i<len(cities); i++{
		if(cities[i].name == s){
			cities[i] = cities[len(cities)-1] 
			cities = cities[:len(cities)-1]
		}
	}
	return cities
}
//prints out what is remaining of the city map
func reconstruct(cities []Cities){
	for _,v := range cities{
		fmt.Print(v.name)
		if v.north != ""{
			fmt.Print(fmt.Sprint(" north=", v.north))
		}
		if v.south != ""{
			fmt.Print(fmt.Sprint(" south=", v.south))
		}
		if v.east != ""{
			fmt.Print(fmt.Sprint(" east=", v.east))
		}
		if v.west != ""{
			fmt.Print(fmt.Sprint(" west=", v.west))
		}
		fmt.Println()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}