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
	fmt.Println(bubbleSort(ary, n))
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

func bubbleSort(ary []int, n int) int {
	flag := true
	sw := 0
	for i := 0; flag; i++ {
		flag = false
		for j := n - 1; i+1 <= j; j-- {
			if ary[j] < ary[j-1] {
				ary[j], ary[j-1] = ary[j-1], ary[j]
				sw++
				flag = true
			}
		}
	}
	print(ary, n)
	return sw
}
