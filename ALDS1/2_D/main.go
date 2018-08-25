package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cnt int

// シェルソート
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	ary := make([]int, n)
	for i := 0; i < n; i++ {
		ary[i] = nextInt(sc)
	}
	shellSort(ary, n)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func insertionSort(ary []int, n, g int) {
	for i := g; i < n; i++ {
		v := ary[i]
		j := i - g
		for 0 <= j && v < ary[j] {
			ary[j+g] = ary[j]
			j -= g
			cnt++
		}
		ary[j+g] = v
	}
}

func shellSort(ary []int, n int) {
	var g []int
	h := 1
	for h <= n {
		g = append(g, h)
		h = 3*h + 1
	}

	glen := len(g)
	fmt.Println(glen)

	var s string
	for i := glen - 1; 0 <= i; i-- {
		s += " " + strconv.Itoa(g[i])
		insertionSort(ary, n, g[i])
	}
	fmt.Println(strings.TrimSpace(s))

	fmt.Println(cnt)
	for _, v := range ary {
		fmt.Println(v)
	}
}
