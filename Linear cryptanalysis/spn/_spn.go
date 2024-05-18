package spn

var S_b_map = map[int]int{
	0: 5,
	1: 3,
	2: 4,
	3: 6,
	4: 0,
	5: 2,
	6: 1,
	7: 7,
}

func xor(m, k []int) []int {
	res := make([]int, len(m))
	for i := 0; i < len(m); i++ {
		res[i] = m[i] ^ k[i]
	}
	return res
}

func sboxApply(input []int) {
	for i, v := range input {
		input[i] = S_b_map[v]
	}
}

func permutation(input []int) []int {
	num1 := (input[0] & 4) | ((input[1] & 4) >> 1) | (input[2] & 4 >> 2)
	num2 := ((input[0] & 2) << 1) | (input[1] & 2) | (input[2] & 2 >> 1)
	num3 := ((input[0] & 1) << 2) | ((input[1] & 1) << 1) | (input[2] & 1)
	return []int{num1, num2, num3}
}

func round(m, k []int, final bool) []int {
	tmp := xor(m, k)
	sboxApply(tmp)
	if final {
		return xor(tmp, k)
	}
	return permutation(tmp)
}

func SPN_net(m, k []int) []int {
	res := m
	for i := 0; i < 2; i++ {
		res = round(res, k, false)
	}
	return round(res, k, true)
}
