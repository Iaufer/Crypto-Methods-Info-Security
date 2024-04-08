// main.go

package main

import (
	"bufio"
	"fmt"
	"la/sdes"
	"os"
	"strconv"
)

func Equal(a1, a2 []uint8) bool {
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func str_to_arr(p []string) ([]uint8, []uint8) {

	var plaintext, dectext []uint8 = []uint8{}, []uint8{}

	for i := 0; i < 8; i++ {
		if p[0][i] == '1' {
			plaintext = append(plaintext, 1)
		} else {
			plaintext = append(plaintext, 0)
		}
		if p[1][i] == '1' {
			dectext = append(dectext, 1)
		} else {
			dectext = append(dectext, 0)
		}
	}
	return plaintext, dectext
}

func main() {

	// plaintextPairs := [][]string{
	//     {"01100101", "00100010"},
	//     {"10111010", "01101110"},
	//     {"11011001", "10001110"},
	//     {"01001011", "00011100"},
	// }

	// _, _ := str_to_arr(plaintextPairs[1])

	// m := []uint8{0, 1, 1, 1, 0, 0, 1, 0}
	// k := []uint8{1, 0, 1, 0, 0, 0, 0, 0, 1, 0}

	// dec := Encrypt(m, k) // false - шифрование, true - деш ифрвание
	// fmt.Println(dec)

	file, err := os.Open("text/plaintext.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	plaintext := make([][]uint8, 0)
	key := []uint8{1, 0, 1, 0, 0, 0, 0, 0, 1, 0}
	
	for scanner.Scan() {
		symbol := scanner.Bytes()
		plain_chunk := str_to_arr_fi(symbol[0])
		plaintext = append(plaintext, plain_chunk)
	}


	if err := scanner.Err(); err != nil{
		fmt.Println("Ошибка при сканировании файла:", err)
	}

	cipher := Encrypt(plaintext, key)
	fmt.Println(cipher)

	in_file_encrypt(cipher)

}


func in_file_encrypt(cipher [][]uint8){
	bits := Decrypt(cipher[0], []uint8{1, 0, 1, 0, 0, 0, 0, 0, 1, 0})

    // Преобразование среза битов в строку
    var bitStr string
    for _, bit := range bits {
        bitStr += strconv.Itoa(int(bit))
    }

    // Преобразование строки битов в целое число
    num, _ := strconv.ParseUint(bitStr, 2, 8) // 8 бит для ASCII

    // Преобразование числа в символ ASCII
    asciiChar := rune(num)

    // Вывод символа ASCII
    fmt.Println(string(asciiChar))
}


func Encrypt(plaintext [][]uint8, key []uint8) [][]uint8 {
	cipher := make([][]uint8, 0)

	for i := 0; i < len(plaintext); i++{
		res := sdes.DES(plaintext[i], key, false)
		cipher = append(cipher, res)
	}

	return cipher
}

func Decrypt(plaintext, key []uint8) []uint8 {
	res := sdes.DES(plaintext, key, true)
	return res
}

func str_to_arr_fi(s byte, ) []uint8 {
	mask := uint8(64)
	tmp := []uint8{}

	for i := 0; i < 8; i++{
		if s&mask != 0{
			tmp = append(tmp, 1)
		}else{
			tmp = append(tmp, 0)
		}
		mask = mask >> 1
		
	}
	return tmp
}
