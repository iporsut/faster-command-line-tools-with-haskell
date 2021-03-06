package main

import (
	"bufio"
	"fmt"
	"os"
)

func processLine(b []byte) (int, int) {
	key := 0
	val := 0
	i := 0
	for b[i] != '\t' {
		i++
	}
	for i++; b[i] != '\t'; i++ {
		key = key*10 + int(b[i]) - '0'
	}
	for i++; b[i] != '\t'; i++ {
		val = val*10 + int(b[i]) - '0'
	}
	return key, val
}

func processFile(file *os.File) (int, int) {
	file.Seek(0, 0)
	var sumByKey [2009]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		k1, v1 := processLine(line)
		sumByKey[k1] += v1
	}
	var k int
	var v int
	for i, val := range sumByKey {
		if val > v {
			k = i
			v = val
		}
	}
	return k, v

}

func main() {
	file, err := os.Open("../ngrams.tsv")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	k, v := processFile(file)
	if k != 2006 || v != 22569013 {
		fmt.Println("bad result")
		// b.Fatalf(`bad result %v | %v`, k, v)
	}
}
