package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ミンコフスキー距離
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	x := make([]float64, n)
	y := make([]float64, n)
	sc.Scan()
	for i, v := range strings.Split(sc.Text(), " ") {
		num, _ := strconv.ParseFloat(v, 64)
		x[i] = num
	}
	sc.Scan()
	for i, v := range strings.Split(sc.Text(), " ") {
		num, _ := strconv.ParseFloat(v, 64)
		y[i] = num
	}
	s := ""
	for i := 1; i <= 3; i++ {
		p := float64(i)
		var D float64
		for j := 0; j < n; j++ {
			D += math.Pow(math.Abs(x[j]-y[j]), p)
		}
		s += fmt.Sprintf("%f\n", math.Pow(D, 1/p))
	}
	var D float64
	for i := 0; i < n; i++ {
		D = math.Max(D, math.Abs(x[i]-y[i]))
	}
	s += fmt.Sprintf("%f\n", D)
	fmt.Print(s)
}
