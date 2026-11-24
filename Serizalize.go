package main

import (
	"bytes"
	"encoding/gob"
)

func (b *Block) Serialize() []byte { //нужно сделать из блока последовательность байтов для сохранения в БД
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b) //сама сериаизция
	if err != nil {
		return nil
	}
	return result.Bytes()
}

func Deserialize(d []byte) *Block { //десериализация-из среза юайт получиь указатель на блок
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d)) //создание потока чтения из среза байт
	err := decoder.Decode(&block)                 //десериализация,заполненеи структкры
	if err != nil {
		return nil
	}
	return &block
}
