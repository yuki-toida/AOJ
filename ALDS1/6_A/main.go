package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var A []int

// パーティション
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	A = make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
	}

	q := partition(0, n-1)

	for i := 0; i < n; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		if i == q {
			fmt.Print("[")
		}
		fmt.Printf("%d", A[i])
		if i == q {
			fmt.Print("]")
		}
		if i == n-1 {
			fmt.Println("")
		}
	}
}

func partition(p, r int) int {
	var x, i, t int
	x = A[r]
	i = p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			t = A[i]
			A[i] = A[j]
			A[j] = t
		}
	}
	t = A[i+1]
	A[i+1] = A[r]
	A[r] = t
	return i + 1
}
