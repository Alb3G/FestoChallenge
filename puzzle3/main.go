package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 10: armed live flipped flipped ready reversed (This trap is safe)

// inactive, disabled, quiet, standby, idle (the trap is deactivated)
// live, armed, ready, primed, active (trap activated)
// flipped, toggled, reversed, inverted, switched (!status)

func getDigit(str string) int {
	trimLine := strings.Split(str, ":")
	res, _ := strconv.Atoi(trimLine[0])
	return res
}

func main() {
	file, err := os.Open("03_trap_logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	trapStatus := false
	id_sum := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		trimLine := strings.Fields(sc.Text())
		for _, str := range trimLine {
			if str == "inactive" || str == "disabled" || str == "quiet" || str == "standby" ||
				str == "idle" {
				trapStatus = true
			} else if str == "live" || str == "armed" || str == "ready" || str == "primed" ||
				str == "active" {
				trapStatus = false
			} else {
				trapStatus = !trapStatus
			}
		}
		if trapStatus {
			fmt.Println("Safe trap")
			id_sum += getDigit(trimLine[0])
		}
	}
	fmt.Println(id_sum)
}
