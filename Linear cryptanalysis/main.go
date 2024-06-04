package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var S_b_map = map[int]int{
	0: 6,
	1: 7,
	2: 4,
	3: 3,
	4: 2,
	5: 5,
	6: 1,
	7: 0,
}

func xor(m, k []int) (res []int) {
	for i := 0; i < len(m); i++ {
		res = append(res, (m[i] ^ k[i]))
	}
	return res
}

func round(m, k []int) []int {
	tmp := xor(m, k)
	for i := 0; i < len(tmp); i++ {
		tmp[i] = S_b_map[tmp[i]]
	}

	num1 := (tmp[0] & 4) | ((tmp[1] & 4) >> 1) | (tmp[2] & 4 >> 2)
	num2 := ((tmp[0] & 2) << 1) | (tmp[1] & 2) | (tmp[2] & 2 >> 1)
	num3 := ((tmp[0] & 1) << 2) | ((tmp[1] & 1) << 1) | (tmp[2] & 1)

	return []int{num1, num2, num3}
}

func last_round(m, k []int) []int {
	tmp := xor(m, k)
	for i := 0; i < len(tmp); i++ {
		tmp[i] = S_b_map[tmp[i]]
	}
	tmp = xor(tmp, k)

	return tmp
}

func SPN_net(m, k []int) []int {
	//первый раунд
	res1 := round(m, k)
	res1 = round(res1, k)
	res1 = last_round(res1, k)
	// for i := 0; i < 3; i++{

	return res1
}

func bin_to_dec(b string) int {
	decimal, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return -1
	}
	return int(decimal)
}

func print_res(res []int) {
	for i := 0; i < len(res); i++ {
		mask := 4
		for j := 0; j < 3; j++ {
			if res[i]&mask != 0 {
				fmt.Print(1)
			} else {
				fmt.Print(0)

			}
			mask >>= 1
		}
		fmt.Print(" ")
	}
	fmt.Println()
}


func analysis() {
	x_list := make([][]int, 0)
	y_list := make([][]int, 0)

	for i := 0; i < 1000; i++{
		x1 := rand.Intn(8)
		x2 := rand.Intn(8)
		x3 := rand.Intn(8)

		x := []int{x1, x2, x3}

		x_list = append(x_list, x)
		y_list = append(y_list, SPN_net(x, []int{bin_to_dec("001"), bin_to_dec("010"), bin_to_dec("011")}))
	}

	// for _, row := range x_list {
	// 	fmt.Println(row)
	// }

	calc(x_list, y_list)
}

func calc(x, y [][]int){
	count := 0
	for i := 0; i < len(x); i++{
		// xy := (y[i][0] & 4) ^ (y[i][0]&2) ^ (x[i][0]&4) ^ (x[i][0]&2)  //первое уравнение

		// xy := (y[i][1] & 4) ^ (y[i][1] & 2) ^ (x[i][1] & 4) ^ (x[i][1] & 2) //второе уравнение

		// xy := (y[i][2] & 4) ^ (y[i][2] & 2) ^ (x[i][2] & 4) ^ (x[i][2] & 2) //третье уравнение
		
		// xy := (y[i][2] & 2) ^ (y[i][2] & 1) ^ (x[i][2] & 1) // шестое уравнение

		xy := (y[i][0]&2) ^ (x[i][0]&4) // пятое уравнение

		if xy == 0{
			count = count + 1
		}
	}

	fmt.Println(count)
}

func main() {
	// m := []int{bin_to_dec("100"), bin_to_dec("101"), bin_to_dec("110")}
	// k := []int{bin_to_dec("001"), bin_to_dec("010"), bin_to_dec("011")} // верный

	analysis()

}
