package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("/rootfs/proc/1/mounts")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		words := strings.Split(line, " ")
		splitWord := strings.Split(words[3], ",")
		var upperLayer string
		if strings.Contains(splitWord[3], "upperlayer") {
			upperLayer = splitWord[3]
		}
		fmt.Println(upperLayer)
	}
}
