package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	ary := make([]int, n)
	for i := 0; i < n; i++ {
		ary[i] = nextInt(sc)
	}
	print(ary, n)
	insertionSort(ary, n)
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

func insertionSort(ary []int, n int) {
	for i := 1; i < n; i++ {
		v := ary[i]
		j := i - 1
		for 0 <= j && v < ary[j] {
			ary[j+1] = ary[j]
			j--
		}
		ary[j+1] = v
		print(ary, n)
	}
}
