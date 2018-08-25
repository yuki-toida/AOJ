package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ベクトルと行列の積
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nm := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])
	A := make([][]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		sc.Scan()
		ary := strings.Split(sc.Text(), " ")
		val := make([]int, m)
		for j := 0; j < m; j++ {
			num, _ := strconv.Atoi(ary[j])
			val[j] = num
		}
		A[i] = val
	}
	for i := 0; i < m; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		b[i] = num
	}

	s := ""
	for _, ary := range A {
		sum := 0
		for j, v := range ary {
			sum += v * b[j]
		}
		s += fmt.Sprintf("%v\n", sum)
	}
	fmt.Print(s)
}
