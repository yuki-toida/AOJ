package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Circle
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a, _ := strconv.ParseFloat(sc.Text(), 64)
	fmt.Printf("%f %f\n", a*a*math.Pi, 2*a*math.Pi)
}
