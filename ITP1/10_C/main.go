package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 標準偏差
func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := ""
	for {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())
		if n == 0 {
			break
		}

		ary := make([]float64, n)
		len := float64(n)
		var sum float64

		sc.Scan()
		for i, v := range strings.Split(sc.Text(), " ") {
			num, _ := strconv.ParseFloat(v, 64)
			ary[i] = num
			sum += num
		}
		m := sum / len
		var val float64
		for _, v := range ary {
			val += math.Pow(v-m, 2)
		}
		s += fmt.Sprintf("%v\n", math.Sqrt(val/len))
	}
	fmt.Print(s)
}
