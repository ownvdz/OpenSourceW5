package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	fmt.Println(inputScore)
}