package main

import (
	"log"

	"go.etcd.io/bbolt"
)

type BlockchainIterator struct { //нужен чтобы пройти от конца цепочки назад (к началу ) по хешу предыдущего блока
	currentHash []byte    //хэш текущего блока
	db          *bbolt.DB //указатель на базу данных
}

func (bc *BlockChain) Iterator() *BlockchainIterator {
	bci := BlockchainIterator{bc.tip, bc.db} //bc.tip-конец (последний хеш) цепочки
	return &bci
}
func (iter *BlockchainIterator) Next() *Block { //чтение блоков от последнего к первому
	var block *Block
	err := iter.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(iter.currentHash) //получаем блок как есть в виде байт после сериализации (gob)
		if encodedBlock == nil {                //если блок не найден
			return nil
		}
		block = Deserialize(encodedBlock) //и десереиализуем
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	iter.currentHash = block.PrevBlockHash //в итератор записываем хеш предыдущего блока
	return block
}
