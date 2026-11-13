package main

import (
	"fmt"
	"strconv"
)

func main() { //just test
	bc := NewBlockChain()
	fmt.Println(bc.blocks)
	bc.AddToBlockChain("Send 1 crypto")
	bc.AddToBlockChain("Send 0.25 crypto")
	for _, block := range bc.blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
		pow := NewProofOFWork(block)
		fmt.Printf("Proof Of Work: %s\n", strconv.FormatBool(pow.Validate()))
	}
}
