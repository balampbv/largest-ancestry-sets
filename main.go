package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//fetch inputs from cmd
	//process each commands
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		inputArr := strings.Split(text, " ")
		switch inputArr[0] {
		case "extract_block":
			blockID := inputArr[1]
			fmt.Println(blockID)
		}

	}
}
