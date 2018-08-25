package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 選択ソート
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	ary := make([]int, n)
	for i := 0; i < n; i++ {
		ary[i] = nextInt(sc)
	}
	fmt.Println(selectionSort(ary, n))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func print(ary []int, n int) {
	s := strconv.Itoa(ary[0])
	for i := 1; i < n; i++ {
		s += " " + strconv.Itoa(ary[i])
	}
	fmt.Println(s)
}

func selectionSort(ary []int, n int) int {
	sw := 0
	for i := 0; i < n-1; i++ {
		minj := i
		for j := i; j < n; j++ {
			if ary[j] < ary[minj] {
				minj = j
			}
		}
		if i < minj {
			ary[i], ary[minj] = ary[minj], ary[i]
			sw++
		}
	}

	print(ary, n)
	return sw
}
