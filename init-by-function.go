package data

import "fmt"

type Point struct {
	x float64
	y float64
}

func (p *Point) initMe(x1, y1 float64) {
	p.x = x1
	p.y = y1
}

func (p *Point) Scale(v float64) {
	p.x = v * p.x
	p.y = v * p.y
}

func (p *Point) PrintP() {
	fmt.Println(p.x, p.y)
}
