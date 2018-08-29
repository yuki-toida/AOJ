package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 線形探索
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	S := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		S[i] = v
	}
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	T := make([]int, q)
	for i := 0; i < q; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		T[i] = v
	}

	cnt := 0
	m := make(map[int]bool, n)
	for _, t := range T {
		for _, s := range S {
			if t == s {
				if _, ok := m[s]; !ok {
					cnt++
					m[s] = true
				}
				continue
			}
		}
	}
	fmt.Println(cnt)
}
