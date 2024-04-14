package main

import(
	"fmt"
	"strconv"
)

var S_b_map = map[int]int{
	0: 7,
	1: 0,
	2: 6,
	3: 5,
	4: 2,
	5: 1,
	6: 3,
	7: 4,
}



func xor(m, k []int)(res []int){
	for i := 0; i < len(m); i++{
		res = append(res, (m[i]^k[i]))
	}
	return res
}

func round(m, k []int)([]int){
	tmp := xor(m, k)
	for i := 0; i < len(tmp); i++{
		tmp[i] = S_b_map[tmp[i]]
	}

	num1 := ((tmp[0]&4)) | ((tmp[1]&4) >> 1) | (tmp[2]&4 >> 2)
	num2 := ((tmp[0]&2) << 1) | ((tmp[1]&2)) | (tmp[2]&2 >> 1)
	num3 := ((tmp[0]&1) << 2) | ((tmp[1]&1) << 1) | (tmp[2]&1)

	
	return []int{num1, num2, num3}
}


func last_round(m, k []int)([]int){
	tmp := xor(m, k)
	for i := 0; i < len(tmp); i++{
		tmp[i] = S_b_map[tmp[i]]
	}
	tmp = xor(tmp, k)
	
	return tmp
}


func SPN_net(m, k []int)([]int){
	//первый раунд
	res1 := round(m, k)
	res1 = round(res1, k)
	res1 = last_round(res1, k)
	// for i := 0; i < 3; i++{

	return res1
}

func bin_to_dec(b string)(int){
	decimal, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return -1
	}
	return int(decimal)
}

func print_res(res []int){
	for i := 0; i < len(res); i++{
		mask := 4
		for j := 0; j < 3; j++{
			if res[i]&mask != 0{
				fmt.Print(1)
			}else {
				fmt.Print(0)
				
			}
			mask >>= 1
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

func main(){

	// m := []int{bin_to_dec("110"), bin_to_dec("011"), bin_to_dec("110")} // верный
	// m := []int{bin_to_dec("111"), bin_to_dec("111"), bin_to_dec("011")} // второй блок не верный
	// m := []int{bin_to_dec("000"), bin_to_dec("111"), bin_to_dec("010")} // второй блок не верный
	// m := []int{bin_to_dec("011"), bin_to_dec("100"), bin_to_dec("000")} // верный
	m := []int{bin_to_dec("100"), bin_to_dec("111"), bin_to_dec("010")} // второй блок не верный // теперь все примеры правильные




	// k := []int{bin_to_dec("111"), bin_to_dec("010"), bin_to_dec("001")}
	k := []int{bin_to_dec("111"), bin_to_dec("010"), bin_to_dec("001")}

	
	res := SPN_net(m, k)
	fmt.Println(res)
	print_res(res)

	table := build_table()

	for _, arr := range table{
		fmt.Println(arr)
	}

}




func build_table()(table [8][8]int){
	for i := 0; i < 8; i++{
		for j := 0; j < 8; j++{
			tmp := i^j
			s1 := S_b_map[i]
			s2 := S_b_map[j]
			c := s1^s2
			table[tmp][c] += 1
		}
	}
	return table
}