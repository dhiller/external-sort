package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"testing"
)

func TestSortOutputFileSorted(t *testing.T) {
	createTestData()
	removeFile("testdata/output")

	doSort()

	file, err := os.Open("testdata/output")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	previousInteger := int64(-1)
	index := 0
	for scanner.Scan() {
		integer, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		if previousInteger != -1 && previousInteger > integer {
			t.Fatal(fmt.Sprintf("line %d: %d > %d", index, previousInteger, integer))
		}
		previousInteger = integer
		index++
	}
}

func TestSortSysMemBelowLimit(t *testing.T) {
	createTestData()
	removeFile("testdata/output")

	doSort()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	sysMiB := int(m.Sys / 1024 / 1024)
	maxMiB := 500
	if sysMiB >= maxMiB {
		t.Fatal(fmt.Sprintf("memory allocation (%d) above limit (%d)", sysMiB, maxMiB))
	}
}

func createTestData() {
	removeFile("testdata/input")
	f, err := os.Create("testdata/input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	minFileSize := int64(math.Pow(1024, 3))
	for i := 0; int64(i) < minFileSize; i++ {
		if i % 100000 == 0 {
			// check size of file
			fileInfo, err := f.Stat()
			if err != nil {
				panic(err)
			}
			if fileInfo.Size() > minFileSize {
				return
			}
		}
		f.WriteString(fmt.Sprintf("%d\n", rand.Int()))
	}
}

func removeFile(fileToRemove string) {
	_, err := os.Stat(fileToRemove)
	if !os.IsNotExist(err) {
		os.Remove(fileToRemove)
	}
	if err != nil {
		panic(err)
	}
}
