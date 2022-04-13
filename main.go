package main

import (
	"bitgo/ancestors"
	"bitgo/extractor"
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
		case "extract_block": //extract_block 000000000000000000076c036ff5119e5a5a74df77abf64203473364509f7732
			blockID := inputArr[1]
			extractor.Extractor(blockID)
		case "find_ancestors": //find_ancestors 000000000000000000076c036ff5119e5a5a74df77abf64203473364509f7732
			blockID := inputArr[1]
			ancestors.FindAncestors(blockID)
		}

	}
}
