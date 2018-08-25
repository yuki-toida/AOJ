package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Sorting three number
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
	sort.SliceStable(ary, func(i, j int) bool {
		v1, _ := strconv.Atoi(ary[i])
		v2, _ := strconv.Atoi(ary[j])
		return v1 < v2
	})
	fmt.Println(strings.Join(ary, " "))
}
