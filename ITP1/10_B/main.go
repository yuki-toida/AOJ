package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 三角形
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(sc.Text(), " ")
	a, _ := strconv.ParseFloat(ary[0], 64)
	b, _ := strconv.ParseFloat(ary[1], 64)
	c, _ := strconv.ParseFloat(ary[2], 64)
	sin := math.Sin(c * math.Pi / 180)
	cos := math.Cos(c * math.Pi / 180)
	fmt.Printf("%f\n", a*b*sin/2)
	fmt.Printf("%f\n", a+b+(math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)-2*a*b*cos)))
	fmt.Printf("%f\n", b*sin)
}
