package main

import (
	"io"
	"os"
	"bufio"
	"fmt"
)

func main() {
	f, err := os.Open("test.txt")  //change the file name
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buf))
	}
}