package bdfs

/*
	solution 1: use array
	solution 2: sort
*/

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	res := make([][]string, 0)
	// array 可以作为 key，slice 不行
	m := make(map[[26]int][]string)

	for _, a := range strs {
		key := [26]int{}
		for _, c := range a {
			key[c-'a']++
		}

		m[key] = append(m[key], a)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
