package main

type ans struct {
	res [][]int
}

// func swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (a ans) recurse(soFar []int, toGo []int) {
	if len(toGo) == 0 {
		soFarCopy := append([]int(nil), soFar...)
		a.res = append(a.res, soFarCopy)
	}
	for i := 0; i < len(toGo); i++ {
		toGoCopy := append([]int(nil), toGo...)
		soFarCopy := append(soFar, toGoCopy[i])
		toGoCopy = toGoCopy[0:i]
		if i+1 < len(toGo) {
			toGoCopy = append(toGoCopy, toGoCopy[i+1:]...)
		}
		a.recurse(soFarCopy, toGoCopy)
	}
}

func permute(nums []int) [][]int {
	A := ans{}
	A.recurse([]int{}, nums)
	return A.res
}

func main() {
	permute([]int{1, 2})
}
