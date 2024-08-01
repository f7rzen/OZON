package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var arrLength int
		fmt.Fscan(in, &arrLength)
		arr := make([]int, arrLength)
		for j := 0; j < arrLength; j++ {
			fmt.Fscan(in, &arr[j])
		}

		longestSliceLength := 1

		for i := 0; i < arrLength; i++ {
			for j := i + 1; j < arrLength; j++ {
				slice := arr[i : j+1]
				if hasTwoDigits(slice) {
					if len(slice) > longestSliceLength {
						longestSliceLength = len(slice)
					}
				}
			}
		}

		fmt.Fprintln(out, longestSliceLength)
	}
}

func hasTwoDigits(slice []int) bool {
	digitCount := make(map[int]bool)
	for _, v := range slice {
		digitCount[v] = true
		if len(digitCount) > 2 {
			return false
		}
	}
	return true
}
