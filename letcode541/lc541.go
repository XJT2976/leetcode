package letcode541

func ReverseStr(s string, k int) string {
	length := len(s)
	ret := []byte(s)
	start := 0
	end := 0

	for start < length {

		if start+k-1 >= length {
			end = length - 1
		} else {
			end = start + k - 1
		}

		for begin := start; begin < end; {
			ret[begin], ret[end] = ret[end], ret[begin]
			begin++
			end--
		}

		start += 2 * k
	}

	return string(ret)
}
