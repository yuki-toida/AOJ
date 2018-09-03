package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n, q int
var A []int

// 総当たり
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	A = make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
	}

	sc.Scan()
	q, _ = strconv.Atoi(sc.Text())
	for i := 0; i < q; i++ {
		sc.Scan()
		m, _ := strconv.Atoi(sc.Text())
		if isSolve(0, m) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func isSolve(i, m int) bool {
	if m == 0 {
		return true
	}
	if n <= i {
		return false
	}
	return isSolve(i+1, m) || isSolve(i+1, m-A[i])
}
