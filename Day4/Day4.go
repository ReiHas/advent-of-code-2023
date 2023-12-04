package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func contains(list []string, s string) bool {
	for _, v := range list {
		if s == "" || v == "" {
			continue
		}
		if s == v {
			return true
		}
	}
	return false
}

func main() {
	f, _ := os.Open("input.txt")
	bu := bufio.NewReaderSize(f, 1024)

	sum := 0
	for {
		line, _, err := bu.ReadLine()
		if err == io.EOF {
			break
		}

		card := strings.Split(string(line), ":")
		content := strings.Split(card[1], "|")

		winningNumbers := strings.Split(content[0], " ")
		havingNumbers := strings.Split(content[1], " ")

		matchCount := 0
		for _, v := range havingNumbers {
			if contains(winningNumbers, v) {
				if matchCount == 0 {
					matchCount++
				} else {
					matchCount = matchCount * 2
				}
			}
		}

		sum = sum + matchCount
	}

	fmt.Println(strconv.Itoa(sum))
}
