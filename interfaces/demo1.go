package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type rectangle struct {
	width, height float64
}
type circle struct {
	raidus float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return c.raidus * c.raidus * math.Pi
}
func school(s shape) {
	fmt.Println(s.area())
}
func Demo1() {
	r := rectangle{width: 10, height: 6}
	school(r)
}
