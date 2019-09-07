
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	// create struct instance
	v := Vertex{1, 2}
	fmt.Println(v)

	// create pointer to struct instance
	p := &v

	// modify struct field via pointer
	p.X = 3
	fmt.Println(v)

	v1 := Vertex{1, 2}  // has type Vertex
	v2 := Vertex{X: 1}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	p2 := &Vertex{1, 2} // has type *Vertex

	fmt.Println(v1, v2, v3, p2)

}
