package main

import "fmt"

func main() {
	test1 := []string{"hrn", "hrf", "er", "enn", "rfnn"}
	test2 := []string{"z", "o"}
	fmt.Println(
		AlienDictinary(test1),
		AlienDictinary(test2),
	)
}

func AlienDictinary(words []string) string {
	// return their lexographical ordering
	type set map[rune]bool
	type cList map[rune]set
	// we need a mapping of rune -> rune
	// to reprsent an adj list
	wList := make(map[string]cList)
	for _, w := range words {
		wList[w] = make(cList)
		for _, r := range w {
			wList[w][r] = make(set)
		}
	}

	for i := range len(words) - 1 {
		w1, w2 := words[i], words[i+1]
		smallerLen := len(w1)
		if len(w2) < smallerLen {
			smallerLen = len(w2)
		}
	}
	return words[0]
}
