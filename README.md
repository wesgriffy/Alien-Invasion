# Alien Invasion Project
Go implementation of Alien Invasion. 
Simulate aliens invading a city map, when two aliens are in the same city, they destroy the city and kill each other. Aliens can move north, south, east, or west throughout the map but they cannot move through destroyed cities.

## Project Direction
I started the project by making a struct of Aliens, with fields being their name, the current city they were in, and a list of possible cities they could visit. This proved troublesome as it was hard to store the cities and their neighbors while preserving direction.

I then decided to scrap what I had and start over with a struct for each city, with the fields being the city name, bordering cities, and the number of aliens in that city as an integer.

### Assumptions
I operated under the assumptions that:
1) No cities would have spaces in their name.
2) If there were more than two aliens in a city they would all fight and kill each other.
3) The number of aliens is less than or equal to the number of cities on the map.

### The Invasion
I wrote a function `makeMove()` that would take the alien in each city (if there was one) and move it to a random bordering city. After all the aliens moved, I would run the function `checkCities()` which would check to see if there was more than one alien in the city and make them fight, thus destroying the city and themselves.

## How to Run
Enter in the command line:

`go run alieninvasion.go n filename`

Sample output:
```
$ go run alieninvasion.go 5 testfile2.txt
Baz has been destroyed by alien0, alien1, and alien2
Bar has been destroyed by alien3 and alien4
Foo south=Qu-ux
Bee
Qu-ux north=Foo west=Aag
Aag east=Qu-ux
```