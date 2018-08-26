package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 行列の積
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nml := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(nml[0])
	m, _ := strconv.Atoi(nml[1])
	l, _ := strconv.Atoi(nml[2])
	A := make([][]string, n)
	B := make([][]string, m)
	for i := 0; i < n; i++ {
		sc.Scan()
		A[i] = make([]string, m)
		for j, v := range strings.Split(sc.Text(), " ") {
			A[i][j] = v
		}
	}
	for i := 0; i < m; i++ {
		sc.Scan()
		B[i] = make([]string, l)
		for j, v := range strings.Split(sc.Text(), " ") {
			B[i][j] = v
		}
	}

	C := make([][]int64, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int64, l)
		for j := 0; j < l; j++ {
			for k := 0; k < m; k++ {
				v1 := parseInt64(A[i][k])
				v2 := parseInt64(B[k][j])
				C[i][j] += v1 * v2
			}
		}
	}

	s := ""
	for _, x := range C {
		for _, y := range x {
			s += fmt.Sprintf("%v ", y)
		}
		s = strings.TrimSpace(s) + "\n"
	}
	fmt.Print(s)
}

func parseInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}
