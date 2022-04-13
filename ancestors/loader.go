package ancestors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type InTransaction struct {
	TransactionID string `json:"Txid"`
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

	var transactions map[string][]InTransaction
	json.Unmarshal(byteValue, &transactions)
	fmt.Println(transactions["ffd3d66ae8507af903102c5cf3d8a061e0253ebfbae02dc3a7a0603b5adb12d6"])
}
