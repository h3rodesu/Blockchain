package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct { //содеражание одного блока
	TimeStamp     int64 //время
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
}

func (b *Block) SetHash() { //хешировние
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))                       //конвертация времени к строке в десятичном формате и приведение к срезу байт
	headers := bytes.Join([][]byte{b.Data, b.PrevBlockHash, timestamp}, []byte{}) //разаделитетль инициазрован но пуст чтобы не было мусора
	//хеш не парсится и смысла от разделителя нет
	hash := sha256.Sum256(headers) //[32]byte-массив
	b.Hash = hash[:]
}
func NewBlock(data string, prblhash []byte) *Block { //создание блока
	block := &Block{time.Now().Unix(), prblhash, []byte{} /*хэш изначально пуст но надо указать поле */, []byte(data)}
	block.SetHash()
	return block
}

type BlockChain struct {
	blocks []*Block
}

func (chain *BlockChain) AddToBlockChain(data string) { //доблав
	prevblock := chain.blocks[len(chain.blocks)-1] //предыдущий блок(по дефолту последний)
	nb := NewBlock(data, prevblock.Hash)
	chain.blocks = append(chain.blocks, nb)
}

func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
func NewBlockChain() *BlockChain {
	gen := GenesisBlock()
	newChain := &BlockChain{[]*Block{gen}}
	return newChain
}
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
	}
}
