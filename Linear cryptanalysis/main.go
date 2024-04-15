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
	// m := []int{bin_to_dec("100"), bin_to_dec("111"), bin_to_dec("010")} // второй блок не верный // теперь все примеры правильные
	// // k := []int{bin_to_dec("111"), bin_to_dec("010"), bin_to_dec("001")}
	// k := []int{bin_to_dec("111"), bin_to_dec("010"), bin_to_dec("001")}
	// res := SPN_net(m, k)
	// fmt.Println(res)
	// print_res(res)




	A := []int{bin_to_dec("110"), bin_to_dec("000"), bin_to_dec("000")}

	analysis(A)

}


func analysis(A []int){
	table, _ := build_table(), 0
	tmp_A := []int{}

	for _, el := range A{
		for k := 0; k < len(table); k++{
			if el == k{
				for i := 0; i < len(table[k]); i++{
					if table[k][i] == 8{
						tmp_A = append(tmp_A, i)
					}
				}
			}
		}
	}

	//second round
	m_var := []int{}
	for _, el := range tmp_A{
		for k := 0; k < len(table); k++{
			if el == k{
				for i := 0; i < len(table[k]); i++{
					if table[k][i] == 4{
						m_var = append(m_var, i)
						// fmt.Println(i)
						// tmp_A = append(tmp_A, i)
					}else if table[k][i] == 8{
						// fmt.Println(i)
					}
				}
			}
		}
	}

	// print_res(tmp_A)
	for _, arr := range table{
		fmt.Println(arr)
	}
	// fmt.Println(m_var)


	arr1 := []int{m_var[0], A[1], A[2]}
	arr2 := []int{m_var[1], A[1], A[2]}


	arr1_perm := []int{((arr1[0]&4)) | ((arr1[1]&4) >> 1) | (arr1[2]&4 >> 2), ((arr1[0]&2) << 1) | ((arr1[1]&2)) | (arr1[2]&2 >> 1), ((arr1[0]&1) << 2) | ((arr1[1]&1) << 1) | (arr1[2]&1)}
	arr2_perm := []int{((arr2[0]&4)) | ((arr2[1]&4) >> 1) | (arr2[2]&4 >> 2), ((arr2[0]&2) << 1) | ((arr2[1]&2)) | (arr2[2]&2 >> 1), ((arr2[0]&1) << 2) | ((arr2[1]&1) << 1) | (arr2[2]&1)}

	// print_res(arr1_perm)
	// print_res(arr2_perm)
	
	// X := []int{bin_to_dec("100"), bin_to_dec("111"), bin_to_dec("010")}
	Y := []int{bin_to_dec("111"), bin_to_dec("011"), bin_to_dec("110")}
	// XI := []int{bin_to_dec("010"), bin_to_dec("111"), bin_to_dec("010")}
	YI := []int{bin_to_dec("110"), bin_to_dec("011"), bin_to_dec("011")}

	_Y := xor(Y, YI)
	fmt.Println("-----")
	a_arr := []int{}
	for _, el := range _Y{
		for k := 0; k < len(table); k++{
			if el == k{
				for i := 0; i < len(table[k]); i++{
					if table[i][k] == 4{
						a_arr = append(a_arr, i)
						// m_var = append(m_var, i)
						// fmt.Println(i)
						// tmp_A = append(tmp_A, i)
					}else if table[k][i] == 8{
						fmt.Println(k, "8")
					}
				}
			}
		}
	}

	// print_res(_Y)
	print_res(a_arr)
	print_res(arr1_perm)
	print_res(arr2_perm)

	

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
