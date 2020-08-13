package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "Barcode symbologies are no longer required automatically\n so you'll have to require the ones you need."

	scanner := bufio.NewScanner(strings.NewReader(s))

	// scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanRunes)
	// scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
