package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A/B problems
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
	a, _ := strconv.Atoi(ary[0])
	b, _ := strconv.Atoi(ary[1])
	fmt.Printf("%v %v %f\n", int(a/b), a%b, float64(a)/float64(b))
}
