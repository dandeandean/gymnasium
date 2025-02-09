package main

import "fmt"

func main() {
	test0 := []string{"wrt", "wrf", "er", "ett", "rftt"}
	test1 := []string{"hrn", "hrf", "er", "enn", "rfnn"}
	test2 := []string{"z", "o"}
	fmt.Println(
		AlienDictinary(test0),
		AlienDictinary(test1),
		AlienDictinary(test2),
	)
}

func wordCmp(w1, w2 string) int {
	shorterWord, longerWord := &w1, &w2
	if len(w2) < len(w1) {
		shorterWord, longerWord = &w2, &w1
	}
	i := 0
	for i < len(*shorterWord) {
		if w1[i] != w2[i] {
			return i
		}
		i++
	}
	if i < len(*longerWord) {
		return i
	}
	return -1
}

func AlienDictinary(words []string) string {
	order := make(map[byte][]byte)
	for i := range len(words) - 1 {
		w0, w1 := words[i], words[i+1]
		deltaW := wordCmp(w0, w1)
		order[w0[deltaW]] = append(order[w0[deltaW]], w1[deltaW])
	}
	beenTo := make(map[byte]bool)
	res := make([]byte, 0)
	var posOrderDfs func(b byte) bool
	posOrderDfs = func(b byte) bool {
		wasThere, ok := beenTo[b]
		if ok {
			return wasThere
		}
		beenTo[b] = true
		for _, nextByte := range order[b] {
			if posOrderDfs(nextByte) {
				return true
			}
		}
		beenTo[b] = false
		res = append(res, b)
		return false
	}

	for b := range order {
		if posOrderDfs(b) {
			return ""
		}
	}
	result := ""
	for _, b := range res {
		result = string(b) + result
	}
	return result
}
