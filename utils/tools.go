package utils

import "fmt"

func IntToHex(a int64) []byte {
	//перевод в hex строку и приведение к срезу байт
	return []byte(fmt.Sprintf("%x", a))
}
