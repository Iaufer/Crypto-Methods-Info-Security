// main.go

package main

import (
	"fmt"
	"la/sdes"
	"time"
)

func Equal(a1, a2 []uint8) bool {
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func str_to_arr(p string) []uint8 {

	var plaintext []uint8 = []uint8{}

	for i := 0; i < 8; i++ {
		if p[i] == '1' {
			plaintext = append(plaintext, 1)
		} else {
			plaintext = append(plaintext, 0)
		}
	}
	return plaintext
}

type Data struct {
	m  []uint8
	cf []uint8
}

func RES(keys [][]int, arr []Data, i int) ([][]int, []Data, int) {
	n_keys := make([][]int, 0)

	if i == len(arr) {
		return keys, nil, i
	}

	for k := 0; k < len(keys); k++ {
		tmp := Encrypt(arr[i].m, Int_to_uint8(keys[k][0]))
		tmp = Encrypt(tmp, Int_to_uint8(keys[k][1]))
		if Equal(tmp, arr[i].cf) {
			n_keys = append(n_keys, []int{keys[k][0], keys[k][1]})
		}
	}

	return RES(n_keys, arr, i+1)
}

// func Encrypt_in_file(k1, k2 []uint8) {
// 	file, err := os.Open("text/plaintext.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file: ", err)
// 		return
// 	}

// 	defer file.Close()

// 	buffer := make([]byte, 1)
// 	for {
// 		_, err := file.Read(buffer)
// 		if err != nil {
// 			fmt.Println("The file has been read!")
// 			break
// 		}
// 		sym := inttouint88(int(buffer[0]))

// 		// fmt.Println(string(buffer[0]), sym)
// 		tmp := Encrypt(sym, k1)
// 		tmp = Encrypt(tmp, k2)

// 		Symbol_to_file(tmp, "text/encrypt.txt")

// 		// break
// 	}
// }

// func Decrypt_in_file(k2, k1 []uint8) {
// 	file, err := os.Open("text/encrypt.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file: ", err)
// 		return
// 	}

// 	defer file.Close()

// 	buffer := make([]byte, 1)
// 	for {
// 		_, err := file.Read(buffer)
// 		if err != nil {
// 			fmt.Println("The file has been read!")
// 			break
// 		}

// 		sym := inttouint88(int(buffer[0]))

// 		fmt.Println(string(buffer[0]), sym)
// 		tmp := Decrypt(sym, k1)
// 		tmp = Decrypt(tmp, k2)

// 		fmt.Println(tmp)

// 		Symbol_to_file(tmp, "text/decrypt.txt")

// 		// break
// 	}
// }

// func Symbol_to_file(sym []uint8, path string) {
// 	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)

// 	if err != nil {
// 		fmt.Println("Ошибка открытия файла:", err)
// 		return
// 	}

// 	defer file.Close()
// 	str := ""
// 	for i := 0; i < len(sym); i++ {
// 		if sym[i] == 1 {
// 			str += "1"
// 		} else {
// 			str += "0"
// 		}
// 	}
// 	num, err := strconv.ParseInt(str, 2, 8)

// 	_, err = file.WriteString(string(num))

// 	if err != nil {
// 		fmt.Println("Ошибка записи в файл:", err)
// 		return
// 	}

// }

// func inttouint88(n int) (res []uint8) {
// 	mask := 128

// 	for i := 0; i < 8; i++ {
// 		if mask&n != 0 {
// 			res = append(res, 1)
// 		} else {
// 			res = append(res, 0)
// 		}
// 		mask >>= 1
// 	}
// 	return res
// }

func main() {
	k1 := []uint8{1, 0, 1, 0, 0, 0, 0, 0, 1, 0}
	k2 := []uint8{1, 1, 0, 1, 1, 0, 0, 1, 1, 1}


	dataMap := map[string][]uint8{
		//открытый текст : зашифрованный
		"10100101": Encrypt(Encrypt([]uint8{1,0,1,0,0,1,0,1}, k1), k2),
		"01000101": Encrypt(Encrypt([]uint8{0,1,0,0,0,1,0,1}, k1), k2),
		"11001000": Encrypt(Encrypt([]uint8{1,1,0,0,1,0,0,0}, k1), k2),
		"10010111": Encrypt(Encrypt([]uint8{1,0,0,1,0,1,1,1}, k1), k2),
	}

	arr := []Data{}

	for key, value := range dataMap {
		pair := Data{
			m:  str_to_arr(key),
			cf: value,
		}
		arr = append(arr, pair)
	}

	start := time.Now()

	keys, _ := MITM([]uint8{0,0,0,0,0,1,0,1}, Encrypt(Encrypt([]uint8{0,0,0,0,0,1,0,1}, k1), k2))

	res, _, _ := RES(keys, arr, 0)

	el := time.Since(start)

	for i := 0; i < len(res); i++ {
		fmt.Println(Int_to_uint8(res[i][0]))
		fmt.Println(Int_to_uint8(res[i][1]))
		fmt.Println("-------")
	}

	fmt.Println(len(res))

	fmt.Println("Time to work: ", el.Seconds())
	
}

func Int_to_uint8(n int) (res []uint8) {
	mask := 512

	for i := 0; i < 10; i++ {
		if mask&n != 0 {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
		mask >>= 1
	}
	return res
}

func MITM(m, cf []uint8) (keys [][]int, res [][]uint8) {

	for i := 0; i < (1 << 10); i++ {
		res = append(res, Encrypt(m, Int_to_uint8(i)))
	}

	count := 0

	for i := 0; i < (1 << 10); i++ {
		dec := Decrypt(cf, Int_to_uint8(i))
		for j := 0; j < (1 << 10); j++ {
			if Equal(dec, res[j]) {
				count++
				keys = append(keys, []int{j, i})
				// break
			}
		}
	}
	return keys, res
}

func Encrypt(plaintext []uint8, key []uint8) []uint8 {
	res := sdes.DES(plaintext, key, false)
	return res
}

func Decrypt(plaintext, key []uint8) []uint8 {
	res := sdes.DES(plaintext, key, true)
	return res
}
