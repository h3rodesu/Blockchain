package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) Run() {
	cli.ValidateArgs()                                           //проверка на то что аргументы введены
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError) //выйти если введен некоррктный флаг
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block Data") //добавление блока	addblock -data данные
	switch os.Args[1] {
	case "addblock":
		if err := addBlockCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal(err)
		} //для добавления в блок надо распарсить аргументы: флаг и значение
	case "printchain":
		if err := printChainCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal(err)
		}
	default:
		cli.printUsage() //вывести руководмтво по использованию если флаг некорректен или значение пустое
		os.Exit(1)
	}
	if addBlockCmd.Parsed() { //если в свиче сработало добавление блока (это было распаршено из флага который ввел юзер)
		if *addBlockData == "" { //если значения нет
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData) //если все норм-создаем новый блок с данными которые ввел юзер
	}
	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
func (cli *CLI) addBlock(data string) {
	cli.bc.AddToBlockChain(data)
	fmt.Println("Succes!")
}
func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()
	for {
		block := bci.Next() //возвращает след блок
		fmt.Printf("Prev Hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data %x\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
		pow := NewProofOFWork(block)
		fmt.Println("PoW %s\n", strconv.FormatBool(pow.Validate())) //проверить (валидировать) pow блоков и вывести вердикт
		fmt.Println()
		if len(block.PrevBlockHash) == 0 { //еслии дошли до генезис блока
			break
		}
	}
}
func (cli *CLI) printUsage() {
	fmt.Println("Enter addblock and data  for add block")
	fmt.Println("Enter printchain for watch prev blocks")
}
func (cli *CLI) ValidateArgs() {
	if len(os.Args) <2 {//если введено имя команы без аргументов
		cli.printUsage()
		os.Exit(1)
	}
}
