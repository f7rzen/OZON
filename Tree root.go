package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value    int
	children []int
}

func findChildren(arr []int) []Node {
	var nodes []Node
	i := 0
	for i < len(arr) {
		count := arr[i+1]
		children := arr[i+2 : i+count+2]
		value := arr[i]
		nodes = append(nodes, Node{value, children})
		i += count + 2
	}
	return nodes
}

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
		nodes := findChildren(arr)

		for k := 0; k < len(nodes); k++ {
			flag := true
			for j := 0; j < len(nodes); j++ {
				for _, child := range nodes[j].children {
					if nodes[k].value == child {
						flag = false
						break
					}
				}
				if !flag {
					break
				}
			}
			if flag {
				fmt.Fprintln(out, nodes[k].value)
			}
		}
	}
}
