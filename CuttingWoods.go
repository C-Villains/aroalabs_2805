package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	woods := make([]int, n)
	max := 0

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &woods[i])

		if max < woods[i] {
			max = woods[i]
		}
	}

	fmt.Fprintln(w, binSearch(woods, 1, max, m))
}

func binSearch(arr []int, l, r, target int) (result int) {
	mid := (l + r) / 2
	cut := cut(arr, mid)

	if l > r {
		if l == target {
			return l
		}

		return r
	}

	if cut > target {
		result = binSearch(arr, mid+1, r, target)
	} else if cut < target {
		result = binSearch(arr, l, mid-1, target)
	} else {
		result = mid
	}

	return result
}

func cut(arr []int, h int) (result int) {
	result = 0

	for _, v := range arr {
		if v > h {
			result += v - h
		}
	}

	return
}
