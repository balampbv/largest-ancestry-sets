package extractor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BlockInfo struct {
	ID      string `json:"id"`
	TXcount int    `json:"tx_count"`
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
	blockDetails := BlockInfo{}
	json.Unmarshal(body, &blockDetails)

	fmt.Printf("Body : %s", body)
	fmt.Printf("Tx_Count %d", blockDetails.TXcount)

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
