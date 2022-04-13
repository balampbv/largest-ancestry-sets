package ancestors

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TransactionCount struct {
	TransactionID string
	Count         int
}

type TxHeap []TransactionCount

func (h TxHeap) Len() int           { return len(h) }
func (h TxHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h TxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(TransactionCount))
}

func (h *TxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type InTransaction struct {
	TransactionID string `json:"Txid"`
}

func LoadFileToStructs(blockHash string) map[string][]InTransaction {
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

	return transactions
}

func FindAncestors(blockHash string) {
	parentToChildMap := LoadFileToStructs(blockHash)
	childToParentMap := make(map[string][]string)
	for txId, value := range parentToChildMap {
		for _, inTx := range value {
			//fmt.Println(txId, "  ", inTx.TransactionID)
			val := childToParentMap[inTx.TransactionID]
			childToParentMap[inTx.TransactionID] = append(val, txId)
		}
	}

	h := &TxHeap{}
	heap.Init(h)

	for k, v := range childToParentMap {
		h.Push(TransactionCount{
			TransactionID: k,
			Count:         len(v),
		})
	}

	for i := 0; i < 10; i++ {
		tx := h.Pop()
		fmt.Println(i, " => ", tx.(TransactionCount).TransactionID)
	}

	// ancestorMap := make(map[string][]string)
	// for child, parents := range childToParentMap { //fetch all parents for a child
	// 	for _, parent := range parents { //for each parent
	// 		val := ancestorMap[child]
	// 		for _, v := range childToParentMap[parent] { //fetch the parents
	// 			ancestorMap[child] = append(val, v)
	// 		}
	// 	}
	// }

	// for k, v := range ancestorMap {
	// 	fmt.Println(k, " ", v)
	// }

}
