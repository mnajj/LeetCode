func countMatches(items [][]string, k string, v string) int {
	var idx int
	switch k {
	case "color":
		idx = 1
	case "name":
		idx = 2
	}
	var ctr int
	for _, itm := range items {
		if itm[idx] == v {
			ctr++
		}
	}
	return ctr
}
