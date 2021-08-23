package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	doSort()
}

func doSort() {
	file, err := os.Open("testdata/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	integers := []int64{}
	for scanner.Scan() {
		integer, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		integers = append(integers, integer)
	}

	sort.Slice(integers, func(i, j int) bool { return integers[i] < integers[j] })

	output, err := os.Create("testdata/output")
	if err != nil {
		panic(err)
	}
	defer output.Close()
	for _, data := range integers {
		output.WriteString(fmt.Sprintf("%d\n", data))
	}
}
