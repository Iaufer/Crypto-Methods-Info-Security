// main.go

package main

import (
	"fmt"
	"la/sdes"
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

func main() {

	// dataMap := map[string][]uint8{
	// 	//открытый текст : зашифрованный
	// 	// "01100101": []uint8{0, 0, 0, 0, 0, 1, 0, 1},
	// 	"10100101": []uint8{1, 0, 0, 1, 0, 0, 1, 1},
	// 	"01000101": []uint8{0, 1, 1, 1, 0, 0, 1, 1},
	// 	"11001000": []uint8{0, 0, 0, 1, 1, 0, 1, 1},
	// 	"10010111": []uint8{0, 1, 0, 1, 0, 0, 1, 1},
	// }

	dataMap := map[string][]uint8{
		//открытый текст : зашифрованный
		// "01100101": []uint8{0, 0, 0, 0, 0, 1, 0, 1},
		"10011001": []uint8{1, 0, 1, 0, 1, 0, 1, 1},
		"10001010": []uint8{0, 0, 1, 1, 1, 1, 1, 0},
		"01010111": []uint8{0, 1, 0, 0, 0, 1, 0, 1},
	}
	// KEY1 = '1010101010'
	// KEY2 = '1111011011'
	// a := []uint8{1,0,1,0,1,1,1,1}
	k := []uint8{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}
	k1 := []uint8{0, 1, 0, 1, 0, 1, 0, 0, 1, 1}

	// q := Encrypt(a, k)
	// q = Encrypt(q, k1)
	// fmt.Println(q)

	arr := []Data{}

	for key, value := range dataMap {
		pair := Data{
			m:  str_to_arr(key),
			cf: value,
		}
		arr = append(arr, pair)
	}

	//res, _ :=
	// keys, _ := MITM([]uint8{0, 1, 1, 0, 0, 1, 0, 1}, []uint8{0, 0, 0, 0, 0, 1, 0, 1})
	keys, _ := MITM([]uint8{1, 0, 1, 0, 1, 1, 1, 1}, []uint8{1, 0, 1, 1, 1, 1, 0, 1})

	res, _, _ := RES(keys, arr, 0)
	// fmt.Println("-------")
	// fmt.Println(Int_to_uint8(res[0][0]))
	// fmt.Println(Int_to_uint8(res[0][1]))
	// fmt.Println("-------")
	// fmt.Println(Int_to_uint8(res[1][0]))
	// fmt.Println(Int_to_uint8(res[1][1]))

	for i := 0; i < len(res); i++ {
		fmt.Println(Int_to_uint8(res[i][0]))
		fmt.Println(Int_to_uint8(res[i][1]))
		fmt.Println("-------")
	}
	fmt.Println(len(res))

	if Equal(Int_to_uint8(res[0][0]), k) && Equal(Int_to_uint8(res[0][1]), k1) {
		fmt.Println("Succes")
	}

	// k := []uint8{1, 0, 1, 0, 0, 0, 0, 0, 1, 0}
	// k1 := []uint8{1, 1, 0, 1, 1, 0, 0, 1, 1, 1}
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
