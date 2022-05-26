package leetcode71

import "strings"

func SimplifyPath(path string) string {
	ret := make([]string, 0)

	sub := strings.Split(path, "/")
	for _, s := range sub {
		if s == ".." {
			if len(ret) > 0 {
				ret = ret[:len(ret)-1]
			}
		} else if s == "" || s == "." {
			continue
		} else {
			ret = append(ret, s)
		}
	}

	if len(ret) == 0 {
		return "/"
	}

	res := ""
	for _, s := range ret {
		res += "/" + s
	}

	return res
}
