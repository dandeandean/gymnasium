package main

import "fmt"

func initMapping() map[byte][]byte {
	g := make(map[byte][]byte)
	g['0'] = []byte{'+'}
	g['2'] = []byte{'a', 'b', 'c'}
	g['3'] = []byte{'d', 'e', 'f'}
	g['4'] = []byte{'g', 'h', 'i'}
	g['5'] = []byte{'j', 'k', 'l'}
	g['6'] = []byte{'m', 'n', 'o'}
	g['7'] = []byte{'p', 'q', 'r', 's'}
	g['8'] = []byte{'t', 'u', 'v'}
	g['9'] = []byte{'w', 'x', 'y', 'z'}
	return g
}

func letterCombinations(digits string) []string {
	var dfs func(string)
	res := make([]string, 0)
	m := initMapping()
	dfs = func(s string) {
		if len(s) >= len(digits) {
			if len(s) == 0 {
				return
			}
			res = append(res, s)
			return
		}
		for _, c := range m[digits[len(s)]] {
			dfs(s + string(c))
		}
		return

	}
	dfs("")
	return res
}

func main() {
	fmt.Println(
		letterCombinations("23"),
		letterCombinations("565"),
		letterCombinations("2"),
		letterCombinations(""),
	)
}
