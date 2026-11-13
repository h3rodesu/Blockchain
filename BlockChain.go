package main

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
