package main

import (
	"time"
)

type Block struct { //содеражание одного блока
	TimeStamp     int64 //время
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
	Nonce         int
}

func NewBlock(data string, prblhash []byte) *Block { //создание блока
	block := &Block{time.Now().Unix(), prblhash, []byte{} /*хэш изначально пуст но надо указать поле */, []byte(data), 0}
	pow := NewProofOFWork(block) //получение порога
	nonce, hash := pow.Run()     //сюда вернется nonce и намайненый хэш,валидация блока
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
