package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type solution struct {
	key     []rune
	hammers [][]rune
}

func initSolution() solution {
	s := solution{
		key: []rune{},
		hammers: [][]rune{
			{'B', 'C'},
			{'C', 'B'},
			{'D', 'D'},
			{'B', 'D'},
			{'C', 'D'},
			{'F', 'E'},
			{'A', 'F'},
			{'F', 'A'},
		},
	}
	return s
}

func cleanLine(str string) [][]int {
	a := [][]int{}
	numStr := ""
	for _, char := range str {
		if char >= '0' && char <= '9' {
			numStr += string(char)
		} else if numStr != "" {
			num, _ := strconv.Atoi(string(numStr))
			if len(a) == 0 || len(a[len(a)-1]) == 2 {
				a = append(a, []int{num})
			} else {
				a[len(a)-1] = append(a[len(a)-1], num)
			}
			numStr = ""
		}
	}
	return a
}

func main() {
	f, err := os.Open("11_keymaker_recipe.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		fmt.Println(cleanLine(line))
	}

}
