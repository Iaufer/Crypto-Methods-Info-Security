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

	for i := 0; i < 10000; i++ {
		x1 := rand.Intn(8)
		x2 := rand.Intn(8)
		x3 := rand.Intn(8)

		x := []int{x1, x2, x3}

		x_list = append(x_list, x)
		y_list = append(y_list, SPN_net(x, []int{bin_to_dec("111"), bin_to_dec("010"), bin_to_dec("011")}))

		// y_list = append(y_list, SPN_net(x, []int{bin_to_dec("001"), bin_to_dec("010"), bin_to_dec("011")}))
	}

	// for _, row := range x_list {
	// 	fmt.Println(row)
	// }

	calc(x_list, y_list)
}

func six_equ(i int, x, y [][]int) int {
	var y1, y2, x1 int
	if (y[i][2] & 2) != 0 {
		y1 = 1
	} else {
		y1 = 0
	}

	if (y[i][2] & 1) != 0 {
		y2 = 1
	} else {
		y2 = 0
	}

	if (x[i][2] & 1) != 0 {
		x1 = 1
	} else {
		x1 = 0
	}

	return (y1 ^ y2 ^ x1)
}

func five_equ(i int, x, y [][]int) int {
	// xy := (y[i][0]&2) ^ (x[i][0]&4) // пятое уравнение

	var y1, x1 int

	if (y[i][0] & 2) != 0 {
		y1 = 1
	} else {
		y1 = 0
	}

	if (x[i][0] & 4) != 0 {
		x1 = 1
	} else {
		x1 = 0
	}

	return (y1 ^ x1)
}

func four_equ1(i int, x, y [][]int) int {
	var y1, x1 int

	if (y[i][1] & 1) != 0 {
		y1 = 1
	} else {
		y1 = 0
	}

	if (x[i][0] & 1) != 0 {
		x1 = 1
	} else {
		x1 = 0
	}

	return (y1 ^ x1)
}

func four_equ2(i int, x, y [][]int)(int){
	var y1, x1, y2 int
	if (y[i][1] & 1) != 0 {
		y1 = 1
	} else {
		y1 = 0
	}

	if (y[i][1] & 2) != 0{
		y2 = 1
	} else {
		y2 = 0
	}

	if (x[i][0]&1) != 0{
		x1 = 1
	} else {
		x1 = 0
	}

	return (y1 ^ y2 ^ x1)
}

func four_equ3(i int, x, y [][]int) int {
	var y1, x1, y2 int

	if (y[i][1] & 4) != 0 {
		y1 = 1
	} else {
		y1 = 0
	}

	if (y[i][1] & 1) != 0 {
		y2 = 1
	} else {
		y2 = 0
	}

	if (x[i][0] & 1) != 0 {
		x1 = 1
	} else {
		x1 = 0
	}

	return (y1 ^ y2 ^ x1)
}

func three_equ(i int, x, y [][]int) int {
	var y1, y2, x1, x2 int

	if (y[i][2]&4) != 0 {
		y1 = 1
	} else {
		y1 = 0
	}

	if (y[i][2]&2) != 0 {
		y2 = 1
	} else {
		y2 = 0
	}

	if (x[i][2]&4) != 0 {	
		x1 = 1
	} else {
		x1 = 0
	}

	if (x[i][2]&2) != 0 {
		x2 = 1
	} else {
		x2 = 0
	}

	return (y1 ^ y2 ^ x1 ^ x2)
}

func two_equ(i int, x, y [][]int) int {
	var y1, y2, x1, x2 int

	if (y[i][1]&4) != 0 {
		y1 = 1
	} else {
		y1 = 0
	}

	if (y[i][1]&2) != 0 {
		y2 = 1
	} else {
		y2 = 0
	}

	if (x[i][1]&4) != 0 {
		x1 = 1
	} else {
		x1 = 0
	}

	if (x[i][1]&2) != 0 {
		x2 = 1
	} else {
		x2 = 0
	}

	return (y1 ^ y2 ^ x1 ^ x2)
}

func one_equ(i int, x, y [][]int) int {
	var y1, y2, x1, x2 int

	if (y[i][0]&4) != 0 {
		y1 = 1
	} else {
		y1 = 0
	}

	if (y[i][0]&2) != 0 {
		y2 = 1
	} else {
		y2 = 0
	}

	if (x[i][0]&4) != 0 {
		x1 = 1
	} else {
		x1 = 0
	}

	if (x[i][0]&2) != 0 {
		x2 = 1
	} else {
		x2 = 0
	}

	return (y1 ^ y2 ^ x1 ^ x2)
}

func calc(x, y [][]int) {
	count := 0

	for i := 0; i < len(x); i++ {
		//Первое уравнение
		xy := one_equ(i, x, y)

		//второе уравнение
		// xy := two_equ(i, x, y)
		
		//третье уравнение
		// xy := three_equ(i, x, y)
		
		//первое четвертое уравнение
		// xy := four_equ1(i, x, y)
		//второе четвертое уравнение
		// xy := four_equ2(i, x, y)
		//третье четвертое уравнение
		// xy := four_equ3(i, x, y)


		//пятое уравнение
		// xy := five_equ(i, x, y)

		//шестое уравнение
		// xy := six_equ(i, x, y)

		if xy == 0 {
			count = count + 1
		}
	}
	var p, q float64 = 0.25, 0.5


	if count > (10000/2){
		if p > q{
			fmt.Println("Левая часть равна: ", 0)
		}else if p < q{
			fmt.Println("Левая часть равна: ", 1)
		}
	}else if count < (10000/2){
		if p < q{
			fmt.Println("Левая часть равна: ", 0)
		}else if p > q{
			fmt.Println("Левая часть равна: ", 1)
		}
	}

	fmt.Println(count)
}

func main() {
	// m := []int{bin_to_dec("100"), bin_to_dec("101"), bin_to_dec("110")}
	// k := []int{bin_to_dec("001"), bin_to_dec("010"), bin_to_dec("011")} // верный

	analysis()

}
