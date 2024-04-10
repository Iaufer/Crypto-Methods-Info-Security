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
	// m := []uint8{0, 1, 1, 1, 0, 0, 1, 0}
	m := []uint8{0, 1, 1, 0, 0, 1, 0, 1}


	// k := []uint8{1, 0, 1, 0, 0, 0, 0, 0, 1, 0}
	// k1 := []uint8{1, 1, 0, 1, 1, 0, 0, 1, 1, 1}

	// dec1 := Encrypt(m, k) // false - шифрование, true - деш ифрвание
	// fmt.Println("ШИФР 1 ключ: ", dec1)
	// dec2 := Encrypt(dec1, k1)
	// fmt.Println("Зашифр текст: ", dec2)

	// res, res1 := MITM(m, dec2)
	MITM(m, []uint8{0, 0, 1, 0, 0, 0, 1, 0})
	// for i := 0; i < len(res); i++{
	// 	k := Int_to_uint8(i)
	// 	if Equal(k, []uint8{1, 0, 1, 0, 0, 0, 0, 0, 1, 0}){
	// 		fmt.Println(res[i], "| ", k)
	// 	}
	// }

	// for i := 0; i < len(res1); i++{
	// 	k := Int_to_uint8(i)
	// 	if Equal(k, []uint8{1, 1, 0, 1, 1, 0, 0, 1, 1, 1}){
	// 		fmt.Println(res1[i], k)
	// 	}
	// }
	// fmt.Println(res[642], Int_to_uint8(642))
	// fmt.Println(res1[871], Int_to_uint8(871))

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

func MITM(m, cf []uint8) (res [][]uint8, res_1 [][]uint8) {
	for i := 0; i < (1 << 10); i++ {
		res = append(res, Encrypt(m, Int_to_uint8(i)))
	}
	count := 0

	for i := 0; i < (1 << 10); i++ {
		a := Decrypt(cf, Int_to_uint8(i))

		for j := 0; j < len(res); j++ {
			if Equal(a, res[j]) {
				tmp := Encrypt(m, Int_to_uint8(i))
				tmp1 := Encrypt(tmp, Int_to_uint8(j))

				if Equal(cf, tmp1){
					fmt.Println(cf, Int_to_uint8(i))
					fmt.Println(tmp1, Int_to_uint8(j))
					fmt.Println("######")
					count++
				}
			}
		}
		// res_1 = append(res_1, Decrypt(cf, Int_to_uint8(i)))
	}
	fmt.Println("Отраб", count)
	return res, res_1
}

func Encrypt(plaintext []uint8, key []uint8) []uint8 {
	res := sdes.DES(plaintext, key, false)
	return res
}

func Decrypt(plaintext, key []uint8) []uint8 {
	res := sdes.DES(plaintext, key, true)
	return res
}
