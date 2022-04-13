package extractor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const filename = "000000000000000000076c036ff5119e5a5a74df77abf64203473364509f7732.json"

type BlockInfo struct {
	ID      string `json:"id"`
	TXcount int    `json:"tx_count"`
}

type Transaction struct {
	TransactionID  string          `json:"txid"`
	InTransactions []InTransaction `json:"vin"`
}

type InTransaction struct {
	TransactionID string `json:"Txid"`
}

func Extractor(blockHash string) {
	blockInfo := GetBlockDetails(blockHash)
	blockInfo.TXcount = 100

	transactionMap := make(map[string][]InTransaction)
	globalMap := make(map[string][]InTransaction)

	for i := 0; i < blockInfo.TXcount; i += 25 {
		ExtractTransactions(blockHash, i, transactionMap)
		for k, v := range transactionMap {
			globalMap[k] = v
		}
	}

	file, _ := json.MarshalIndent(globalMap, "", " ")

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(string(file)); err != nil {
		panic(err)
	}
}

func GetBlockDetails(blockHash string) BlockInfo {
	blockDetails := BlockInfo{}

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("https://blockstream.info/api/block/" + blockHash)
	if err != nil {
		fmt.Printf("Error %s", err)
		return blockDetails
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return blockDetails
	}
	json.Unmarshal(body, &blockDetails)

	fmt.Printf("Body : %s", body)
	fmt.Printf("Tx_Count %d", blockDetails.TXcount)

	return blockDetails

}

//block 680000 - 000000000000000000076c036ff5119e5a5a74df77abf64203473364509f7732
func ExtractTransactions(blockHash string, offset int, transactionMap map[string][]InTransaction) {
	transactions := []Transaction{}

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("https://blockstream.info/api/block/" + blockHash + "/txs/" + strconv.Itoa(offset))
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

	json.Unmarshal(body, &transactions)

	for i := 0; i < len(transactions); i++ {
		transactionMap[transactions[i].TransactionID] = transactions[i].InTransactions
	}

	fmt.Printf("Body : %v", transactionMap)

}
