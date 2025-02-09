package main

import "fmt"

func main() {
	test0 := []string{"wrt", "wrf", "er", "ett", "rftt"}
	// test1 := []string{"hrn", "hrf", "er", "enn", "rfnn"}
	// test2 := []string{"z", "o"}
	fmt.Println(
		AlienDictinary(test0),
		// AlienDictinary(test1),
		// AlienDictinary(test2),
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
	order := make(map[byte]byte)
	for i := range len(words) - 1 {
		w0, w1 := words[i], words[i+1]
		deltaW := wordCmp(w0, w1)
		fmt.Println(w0, w1, string(w0[deltaW]), "->", string(w1[deltaW]))
		order[w0[deltaW]] = w1[deltaW]
	}
	var posOrderDfs func(i int)
	posOrderDfs = func(i int) {
		return
	}
	return "abc"
}
