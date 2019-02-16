package bit

/*
	solution 1: shift and bit operate
*/

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res*2 + (num>>uint(i))&1
	}
	return res
}
