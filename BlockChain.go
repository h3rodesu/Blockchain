package main

import (
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

const (
	blocksBucket = "blocks" //имя бакета (папки) для хранения блоков
	dbFile       = "blockchain.db"
)

type BlockChain struct {
	//	blocks []*Block блокчейн больше не хрвнит срез из указателей на блоки
	//блоки лежат на диске,а блокчен хранит:
	tip []byte    //хеш последнего блока (конец цепочки)
	db  *bbolt.DB //подключение к файлу БД
}

func (bc *BlockChain) AddToBlockChain(data string) { //доблав
	var lastHash []byte                          //срез байтов -хеш
	err := bc.db.View(func(tx *bbolt.Tx) error { //транзакция только для чтения
		b := tx.Bucket([]byte(blocksBucket)) //получение доступа к папке (бакету с блоками)
		lastHash = b.Get([]byte("l"))        //чтение из базы по ключу l
		return nil                           //
	})
	if err != nil {
		log.Fatal(err)
	}
	newBlock := NewBlock(data, lastHash)          //создаем новый блок учитывай что до этого взяи хещ предыдщуего
	err = bc.db.Update(func(tx *bbolt.Tx) error { //открытие транзакции а добавление нового блока в бд
		b := tx.Bucket([]byte(blocksBucket))
		err = b.Put(newBlock.Hash, newBlock.Serialize()) //запись нового хеша
		if err != nil {
			return err
		}
		err = b.Put([]byte("l"), newBlock.Hash) //обновление ключа
		if err != nil {
			return err
		}
		bc.tip = newBlock.Hash //обновление конца цепочки
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bbolt.Open(dbFile, 0600, nil) //открытие файла в бд
	if err != nil {
		fmt.Errorf("Error of open the database")
		return nil
	}
	err = db.Update(func(tx *bbolt.Tx) error { //транзакция чтение-запись
		b := tx.Bucket([]byte(blocksBucket)) //получение бакета из бд
		if b == nil {                        //если он пуст-создаем новый блокчейн и соответственно новый генезис
			genesis := GenesisBlock()
			b, err = tx.CreateBucket([]byte(blocksBucket)) //создание нового бакета через транзакцию
			err = b.Put(genesis.Hash, genesis.Serialize()) //закидываем хеш генезис блока и генезис в виде байт
			err = b.Put([]byte("l"), genesis.Hash)         //[]byte("l") приведение строки в срез байтов,l-означает крайний блок
			tip = genesis.Hash                             //крайний в цепочке генезис блок
		} else {
			tip = b.Get([]byte("l")) //если блокчейн существет-ищем крайний блок с ключом l (last)
		}
		return nil
	})
	bc := BlockChain{tip, db}
	return &bc
}
