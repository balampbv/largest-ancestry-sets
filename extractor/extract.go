package extractor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BlockInfo struct {
}

func GetBlockDetails(blockHash string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("https://blockstream.info/api/block/" + blockHash)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	fmt.Printf("Body : %s", body)
}

//block 680000 - 000000000000000000076c036ff5119e5a5a74df77abf64203473364509f7732
func ExtractTransactions(blockHash string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("https://blockstream.info/api/block/" + blockHash + "/txs/[25]")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	fmt.Printf("Body : %s", body)

}
