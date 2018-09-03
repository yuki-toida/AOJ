package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cnt int

// マージソート
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	A := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
	}

	sorted := mergeSort(A)
	fmt.Println(strings.Trim(fmt.Sprint(sorted), "[]"))
	fmt.Println(cnt)
}

// Runs mergeSort algorithm on a slice single
func mergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice) / 2
	left := mergeSort(slice[:mid])
	right := mergeSort(slice[mid:])
	return merge(left, right)
}

// Merges left and right slice into newly created slice
func merge(left, right []int) []int {
	size, k, j := len(left)+len(right), 0, 0
	tmp := make([]int, size)

	for i := 0; i < size; i++ {
		if len(left) <= j {
			tmp[i] = right[k]
			k++
		} else if len(right) <= k {
			tmp[i] = left[j]
			j++
		} else if right[k] < left[j] {
			tmp[i] = right[k]
			k++
		} else {
			tmp[i] = left[j]
			j++
		}
		cnt++
	}
	return tmp
}
