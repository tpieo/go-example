package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	fp, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(fp)
	for {
		if !scanner.Scan() {
			break  //scan a line, if EOF, return false
		}
		line := scanner.Text() //get a line
		fmt.Println(line)
	}
}