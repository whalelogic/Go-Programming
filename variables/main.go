package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
)


var (
	a string
	b int
)

type square struct {
	width float32
	height float32 
}

type circle struct {
	radius float32
}

func (c *circle) area() float32 {
	// math.Pi is float64, so we cast it to float32
	return float32(math.Pi) * c.radius * c.radius
}

func (c *circle) perimeter() float32 {
	// For a circle, perimeter IS circumference
	return 2 * float32(math.Pi) * c.radius
}

func (c *circle) circumference() float32 {
	return c.perimeter()
}

type Shape interface {
	area() float32
	perimeter() float32
	circumference() float32
}


func(s *square) area() float32 {
	return s.width * s.height
}

func (s *square) perimeter() float32 {
	return 2 * (s.width + s.height)
}

func (s *square) circumference() float32 {
	return s.perimeter()
}

func printShape(s Shape) {
	fmt.Println("-----Shape Details-----")
	fmt.Printf("Area:          %0.2f\n", s.area())
	fmt.Printf("Perimeter:     %0.2f\n", s.perimeter())
}

func generalArea(w, h float32) float32 {
	return w * h // width * height
}


func main() {

	add := func(x, y float32) float32 {
		return x + y
	}

	result := add(5, 10)
	fmt.Println("Result of add(5, 10):", result) // Result of add(5, 10): 15
	fmt.Println("Type of variable 'a':", reflect.TypeOf(a)) // Type of variable 'a': string
	fmt.Println("Type of variable 'b':", reflect.TypeOf(b)) // Type of variable 'b': int

	a, b = "Keith's Age", 37
	fmt.Println(a, b)

	sq := new(square)
	sq.height = 14.2 
	sq.width = 13.7
	fmt.Println(sq.area())
	printShape(sq)
	// or 
	area := sq.area()
	perimeter := sq.perimeter()
	fmt.Println("Area of square:", area)
	fmt.Printf("Perimeter of square: %+v\n", perimeter)
	fmt.Printf("Current state: %+v\n", sq)
	fmt.Printf("Original state: %+v\n", *sq)
	fmt.Printf("Memory location: %+v\n", &sq)
	printShape(sq)
	// new shape
	shp := new(Shape)
	fmt.Println("new shape: ", shp)

	sq = &square{width: 10, height: 10}

	// 2. Create a Circle
	circ := &circle{radius: 5}

	// 3. Create a Slice of Shapes
	// This slice can hold ANYTHING that implements the Shape interface
	shapes := []Shape{sq, circ}

	// 4. Iterate (Polymorphism in action!)
	fmt.Println("--- Processing Shapes ---")
	
	for i, shape := range shapes {
		fmt.Printf("Shape #%d:\n", i+1)
		fmt.Printf("  Area:          %0.2f\n", shape.area())
		fmt.Printf("  Circumference: %0.2f\n", shape.circumference())
		fmt.Println()
	}

	rand1 := rand.Int()
	rand2 := rand.Int()
	fmt.Println(rand1+rand2)
	fmt.Println(reflect.DeepEqual(rand1, rand2))
	fmt.Println("waahhhh! :/ ")

	// Use the generalArea func 
	fmt.Println(generalArea(4, 12))

}
