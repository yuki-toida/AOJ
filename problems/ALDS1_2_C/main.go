package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type card struct {
	suit  string
	value int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	bs := make([]card, n)
	ss := make([]card, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		rs := []rune(sc.Text())
		suit := string(rs[0])
		value, _ := strconv.Atoi(string(rs[1]))
		bs[i] = card{suit: suit, value: value}
		ss[i] = card{suit: suit, value: value}
	}
	bubbleSort(bs, n)
	print(bs, n, true)

	selectionSort(ss, n)
	stable := true
	for i, v := range bs {
		if v != ss[i] {
			stable = false
			break
		}
	}
	print(ss, n, stable)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func print(ary []card, n int, stable bool) {
	s := ary[0].suit + strconv.Itoa(ary[0].value)
	for i := 1; i < n; i++ {
		s += " " + ary[i].suit + strconv.Itoa(ary[i].value)
	}
	fmt.Println(s)
	if stable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

func bubbleSort(ary []card, n int) {
	flag := true
	for i := 0; flag; i++ {
		flag = false
		for j := n - 1; i+1 <= j; j-- {
			if ary[j].value < ary[j-1].value {
				ary[j], ary[j-1] = ary[j-1], ary[j]
				flag = true
			}
		}
	}
}

func selectionSort(ary []card, n int) {
	for i := 0; i < n-1; i++ {
		minj := i
		for j := i; j < n; j++ {
			if ary[j].value < ary[minj].value {
				minj = j
			}
		}
		if i < minj {
			ary[i], ary[minj] = ary[minj], ary[i]
		}
	}
}
