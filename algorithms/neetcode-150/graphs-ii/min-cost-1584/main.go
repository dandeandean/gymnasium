package main

// An IntHeap is a min-heap of ints.

func minCostConnectPoints(points [][]int) int {
	manDist := func(i, j int) int {
		distX := points[i][0] - points[j][0]
		distY := points[i][1] - points[j][1]
		if distX < 0 {
			distX = -distX
		}
		if distY < 0 {
			distY = -distY
		}
		return distX + distY
	}
	distTo := make([]int, len(points))
	for i := range points {
		distTo[i] = 99999999999999999
	}
	beenTo := make([]bool, len(points))
	cur, edgeCount, res := 0, 0, 0
	for edgeCount < len(points)-1 {
		beenTo[cur] = true
		next := -1

		// all possible connections
		for i := range points {
			if beenTo[i] {
				continue
			}

			d := manDist(cur, i)
			if d < distTo[i] {
				distTo[i] = d
			}

			if next == -1 || distTo[i] < distTo[next] {
				next = i
			}
		}
		res += distTo[next]
		cur = next
		edgeCount++
	}
	return res
}
