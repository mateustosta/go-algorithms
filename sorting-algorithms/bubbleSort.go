package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slc []int, idx int) {
	tmp := slc[idx]
	slc[idx] = slc[idx+1]
	slc[idx+1] = tmp
}

func BubbleSort(slc []int) {
	var i, j int
	for i = 0; i < len(slc) - 1; i++ {
		for j = 0; j < len(slc) - i - 1; j++ {
			if slc[j] > slc[j+1] {
				Swap(slc, j)
			}
		}
 	}
}

func main() {
	var slc = make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		fmt.Println("Enter a number or enter X to stop: ")
		reader := bufio.NewReader(os.Stdin)
		in, _ := reader.ReadString('\n')

		if in == "x\n" || in == "X\n" {
			break
		}

		num, _ := strconv.Atoi(strings.TrimSuffix(in, "\n"))

		slc = append(slc, num)
	}

	fmt.Println("Unsorted Slice: ")

	BubbleSort(slc)

	fmt.Println("Sorted Slice: ")
	fmt.Println(slc)
}
