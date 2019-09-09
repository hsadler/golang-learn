
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func printMap(m map[string]Vertex) {
	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func main() {

	// create a new map
	var m1 map[string]Vertex = make(map[string]Vertex)
	// add a key and value
	m1["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m1["Bell Labs"])

	// create a map with a map literal
	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	// use range of map to print each key and value
	for k, v := range m2 {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	// map literals can be used without type names
	var m3 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	printMap(m3)


	// DEMONSTRATION: map mutations and data access

	m4 := make(map[string]int)

	// add key and value
	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	// replace key with value
	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	// delete key
	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	// access a key and confirm it exists in the map
	v, ok := m4["Answer"]
	fmt.Println("The value:", v, "Present?", ok)


}
