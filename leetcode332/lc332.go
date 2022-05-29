package leetcode332

import "sort"

type itinerary [][]string

func (itn itinerary) Less(i, j int) bool {
	return itn[i][1] < itn[j][1]
}

func (itn itinerary) Swap(i, j int) {
	itn[i], itn[j] = itn[j], itn[i]
}

func (itn itinerary) Len() int {
	return len(itn)
}

func FindItinerary(tickets [][]string) []string {
	path := []string{"JFK"}
	used := make([]bool, len(tickets))
	sort.Sort(itinerary(tickets))
	backtracking(tickets, "JFK", used, &path)

	return path
}

func backtracking(tickets [][]string, departure string, used []bool, path *[]string) bool {
	if len(*path) == len(tickets)+1 {
		return true
	}

	for i := 0; i < len(tickets); i++ {
		if used[i] {
			continue
		}
		if tickets[i][0] == departure {
			*path = append(*path, tickets[i][1])
			used[i] = true
			if backtracking(tickets, tickets[i][1], used, path) {
				return true
			}
			*path = (*path)[:len(*path)-1]
			used[i] = false
		}
	}

	return false
}
