package ancestors

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

	ancestorCount := make(map[string]int)
	for k, _ := range parentToChildMap {
		ancestorCount[k] = findAncestorCount(parentToChildMap, k, 0)
		// fmt.Println(k, " ", ancestorCount[k])
	}

	var result []TransactionCount
	for k, v := range ancestorCount {
		result = append(result, TransactionCount{
			TransactionID: k,
			Count:         v,
		})
	}

	//Sort based on the ancestor count
	sort.Slice(result[:], func(i, j int) bool {
		return result[i].Count < result[j].Count
	})

	//Heap to return the last 1o transactions
	h := &TxHeap{}
	heap.Init(h)

	for _, v := range result {
		h.Push(TransactionCount{
			TransactionID: v.TransactionID,
			Count:         v.Count,
		})
	}

	for i := 0; i < 10; i++ {
		tx := h.Pop()
		fmt.Println(i, " => ", tx.(TransactionCount).TransactionID)
	}

}

//findAncestorCount recursive funtctions to find the ancestors
func findAncestorCount(childMap map[string][]InTransaction, txId string, result int) int {
	parents, ok := childMap[txId]
	if !ok {
		return 1
	}
	for i := 0; i < len(parents); i++ {
		result = result + findAncestorCount(childMap, parents[i].TransactionID, result)
	}

	return result
}
