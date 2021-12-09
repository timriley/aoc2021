package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	lineCount := 0
	bitCounts := [12]int{}

	for _, line := range strings.Split(string(data), "\n") {
		lineCount += 1

		for i, num := range strings.Split(line, "") {
			num, _ := strconv.Atoi(num)
			bitCounts[i] += num
		}
	}

	threshold := lineCount / 2

	gammaRate := ""
	epsilonRate := ""

	for _, num := range bitCounts {
		if num > threshold {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gammaRateNum, _ := strconv.ParseInt(strings.Join(strings.Fields(gammaRate), ""), 2, 64)
	epsilonRateNum, _ := strconv.ParseInt(strings.Join(strings.Fields(epsilonRate), ""), 2, 64)

	powerConsumption := gammaRateNum * epsilonRateNum
	fmt.Printf("Power consumption: %v\n", powerConsumption)
}
