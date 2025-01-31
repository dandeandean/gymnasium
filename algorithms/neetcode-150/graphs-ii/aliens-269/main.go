package main

func AlienDictinary(words []string) string {
	// return their lexographical ordering
	//adjList := make(map[byte]map[byte]bool)
	for i := range len(words) - 1 {
		w1, w2 := words[i], words[i+1]
		smallerLen := len(w1)
		if len(w2) < smallerLen {
			smallerLen = len(w2)
		}
	}
	return words[0]
}
