package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SENTINEL = 10000

func countingSort(A []int) []int {
	C := make([]int, SENTINEL+1)
	for _, value := range A {
		C[value]++
	}
	for i := 1; i < len(C); i++ {
		C[i] = C[i] + C[i-1]
	}
	B := make([]int, len(A))
	for j := 0; j < len(A); j++ {
		B[C[A[j]]-1] = A[j]
		C[A[j]]--
	}
	return B
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())
	A := make([]int, n)

	for i := range A {
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
	}

	result := strings.TrimRight(fmt.Sprintf("%+v", countingSort(A))[1:], "]")
	fmt.Println(result)

}
