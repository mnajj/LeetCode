
func groupAnagrams(strs []string) [][]string {
	g := map[string][]string{}
	var so string
	var ret [][]string

	for _, s := range strs {
		so = sortStr(s)
		if _, ok := g[so]; ok {
			g[so] = append(g[so], s)
			continue
		}
		g[so] = []string{
			s,
		}
	}
	for _, v := range g {
		ret = append(ret, v)
	}
	return ret
}

func sortStr(s string) string {
	w := []rune(s)
	sort.Slice(w, func(i, j int) bool {
		return w[i] < w[j]
	})
	return string(w)
}
