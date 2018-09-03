package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type point struct {
	x, y float64
}

var radian = math.Pi * 60 / 180

// コッホ曲線
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	p1 := point{x: 0, y: 0}
	p2 := point{x: 100, y: 0}
	print(p1)
	koch(n, p1, p2)
	print(p2)
}

func koch(n int, a, b point) {
	if n == 0 {
		return
	}

	s := point{x: (2*a.x + b.x) / 3, y: (2*a.y + b.y) / 3}
	t := point{x: (a.x + 2*b.x) / 3, y: (a.y + 2*b.y) / 3}
	u := point{
		x: (t.x-s.x)*math.Cos(radian) - (t.y-s.y)*math.Sin(radian) + s.x,
		y: (t.x-s.x)*math.Sin(radian) + (t.y-s.y)*math.Cos(radian) + s.y,
	}

	koch(n-1, a, s)
	print(s)
	koch(n-1, s, u)
	print(u)
	koch(n-1, u, t)
	print(t)
	koch(n-1, t, b)
}

func print(p point) {
	fmt.Printf("%f %f\n", p.x, p.y)
}
