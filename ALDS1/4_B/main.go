package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var n int
var s []int

// 二分探索
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	s = make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		s[i] = v
	}
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	cnt := 0
	for i := 0; i < q; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		if _, err := binarySearch(v); err == nil {
			cnt++
			continue
		}
	}
	fmt.Println(cnt)
}

func binarySearch(key int) (int, error) {
	left, mid, right := 0, 0, n
	for left < right {
		mid = (left + right) / 2
		if s[mid] == key {
			return mid, nil
		} else if key < s[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return 0, errors.New("Not found")
}
