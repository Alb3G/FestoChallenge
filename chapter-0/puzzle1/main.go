package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkSlice(str string) bool {
	i := 0
	for i < len(str)-1 {
		if str[i] > str[i+1] {
			return false
		}
		i++
	}
	return true
}

func main() {
	file, err := os.OpenFile("01_keymaker_ordered.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		if checkSlice(line) {
			fmt.Println(line)
		}
	}
}
