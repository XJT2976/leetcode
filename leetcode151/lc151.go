package leetcode151

func ReverseWords(s string) string {
	ret := []byte(s)

	for i, j := 0, len(s)-1; i < j; {
		ret[i], ret[j] = ret[j], ret[i]
		i++
		j--
	}

	for wordStart := 0; wordStart < len(s); wordStart++ {
		if ret[wordStart] == ' ' {
			continue
		}
		wordEnd := wordStart
		for wordEnd < len(s) && ret[wordEnd] != ' ' {
			wordEnd++
		}
		wordEnd--

		for i, j := wordStart, wordEnd; i < j; {
			ret[i], ret[j] = ret[j], ret[i]
			i++
			j--
		}
		wordStart = wordEnd
	}

	i := 0
	j := 0
	for j < len(s) && ret[j] == ' ' {
		j++
	}

	for j < len(s) {
		if j > 0 && ret[j] == ' ' && ret[j] == ret[j-1] {
			j++
		} else {
			ret[i] = ret[j]
			i++
			j++
		}
	}
	ret = ret[:i]

	j = len(ret) - 1
	for j >= 0 && ret[j] == ' ' {
		j--
	}
	ret = ret[:j+1]

	return string(ret)
}
