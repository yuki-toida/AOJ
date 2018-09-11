package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 挿入ソート
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	ary := make([]int, n)
	for i := 0; i < n; i++ {
		ary[i] = nextInt(sc)
	}
	print(ary)
	insertionSort(ary, n)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func print(ary []int) {
	fmt.Printf("%v\n", strings.Trim(fmt.Sprint(ary), "[]"))
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
		print(ary)
	}
}
