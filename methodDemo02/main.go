package main

import (
	"fmt"
)

type Circle struct{
	radius float64
}

func (c *Circle) area() float64 {
	c.radius = 3
	return 3.14 * c.radius * c.radius
}


func main() {
	var c Circle
	c.radius = 10
	a := c.area()
	fmt.Println("圆的面积是", a)
	fmt.Println(c.radius )

}