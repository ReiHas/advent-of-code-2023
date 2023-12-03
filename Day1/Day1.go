package main

import(
	"bufio"
	"io"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	f, _ := os.Open("input.txt")
	bu := bufio.NewReaderSize(f, 1024)

	var sum int = 0

	for {
		line, _, err := bu.ReadLine()
		if err == io.EOF {
			break
		}
		
		str := string(line)
		strArray := strings.Split(str, "")


		var count int = 0
		var firstInt string
		for {
			target := strArray[count]

			if "0" <= target && target <= "9" {
				firstInt = target
				break
			} else {
				count++
			}
		}
		
		count = len(strArray) - 1
		var lastInt string
		for {
			target := strArray[count]

			if "0" <= target && target <= "9" {
				lastInt = target
				break
			} else {
				count--
			}
		}

		calibrationStr := firstInt + lastInt
		c, _ := strconv.Atoi(calibrationStr)
		sum = sum + c
	}

	fmt.Println(sum)

}