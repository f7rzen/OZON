package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	var p float64
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &p)
		var x, result float64
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &x)
			y := x * p / 100
			result += y - float64(int(y))
		}
		result = math.Round(result*100) / 100
		fmt.Fprintf(out, "%.2f\n", result)
	}
}
