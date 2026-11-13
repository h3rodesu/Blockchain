package main

import (
	"blockchain_project/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const maxNonce = math.MaxInt64 //максимальное число для int64 на случай если хеш не найден

const targbites = 24      //заголовок блока
type ProofOfWork struct { //доказательство верного вычисления хеша блока
	block  *Block
	target *big.Int //сохраненный порог(хеш) с которым в дальнейшем будет сравниваться хеш
}

// 1)вычислить порог
// 2)Перебор случайных чисел (nonce )в блоке,считать хеши и сравнивать с порогом (если хеш больше порога nonce увеличивается на 1)
func NewProofOFWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targbites)) //побитовый сдвиг на 24 байта, получение порога(числа больше него-не валидны)
	//сдвиг так же нужен чтобы полученный порог был меньше 256 бит
	pow := &ProofOfWork{b, target}
	return pow
}
func (pow *ProofOfWork) DataPrepare(nonce int) []byte { //Подготовка данных блока с учетом nonce,просто закинуть в срез байт
	prepData := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		utils.IntToHex(pow.block.TimeStamp), //sha256 требует hex строку
		utils.IntToHex(int64(nonce)),        //sha256 требует hex строку
		utils.IntToHex(targbites),           //защита от подделки после попадания всеть путем изменения сложности
	}, []byte{})
	return prepData
}
func (pow *ProofOfWork) Run() (int, []byte) { //вернет nonce и сам хеш
	var hashInt big.Int
	var hash [32]byte //возврат от sha256
	nonce := 0
	fmt.Printf("Mining,data: %s \n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.DataPrepare(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\rhash: %x", hash)
		hashInt.SetBytes(hash[:])          //интерпритирует как одно 256-битное число
		if hashInt.Cmp(pow.target) == -1 { //сравнение полученного числа хеша с порогом
			break //если меньше-успех
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}
func (pow *ProofOfWork) Validate() bool { //проверка валидности блока
	var hashInt big.Int
	data := pow.DataPrepare(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	if hashInt.Cmp(pow.target) == -1 {
		return true
	}
	return false
}
