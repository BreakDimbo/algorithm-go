package array

/*
	solutoin 1: two pointers with sliding window
		map+formed
		如何判断当前的 滑动窗口 满足 t
		字符串 substring 的问题，模版：map + sliding window
*/

/*
int findSubstring(string s){
	vector<int> map(128,0);
	int counter; // check whether the substring is valid
	int begin=0, end=0; //two pointers, one point to tail and one  head
	int d; //the length of substring

	for() { initialize the hash map here  }

	while(end<s.size()){

			if(map[s[end++]]-- ?){ modify counter here  }

			while(counter condition) {

					 update d here if finding minimum

					//increase begin to make it invalid/valid again

					if(map[s[begin++]]++ ?){ modify counter here }
			}
			update d here if finding maximum
	}
	return d;
}
*/

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	dict := make(map[byte]int)
	for i := range t {
		dict[t[i]]++
	}
	required := len(dict)

	formed := 0
	curWindowCount := make(map[byte]int)
	l, r := 0, 0
	ans := []int{-1, 0, 0} // window length, left, right

	for r < len(s) {
		curWindowCount[s[r]]++
		if count, ok := dict[s[r]]; ok && count == curWindowCount[s[r]] {
			formed++
		}

		for formed == required {
			// record min window
			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}
			curWindowCount[s[l]]--
			if c, ok := dict[s[l]]; ok && c != curWindowCount[s[l]] {
				formed--
			}
			l++
		}
		r++
	}
	if ans[0] == -1 {
		return ""
	} else {
		return s[ans[1] : ans[2]+1]
	}
}
