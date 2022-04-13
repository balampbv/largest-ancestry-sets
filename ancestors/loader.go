package ancestors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type InTransaction struct {
	TransactionID string `json:"txid"`
}

func LoadFile(blockHash string) {
	// Open our jsonFile
	jsonFile, err := os.Open(blockHash + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Println(string(byteValue))

	var transactions map[string][]InTransaction
	json.Unmarshal(byteValue, &transactions)

	fmt.Println(transactions)
}
