package stackk

// https://leetcode.com/problems/car-fleet/

import "sort"

func CarFleet(target int, position []int, speed []int) int {
	ETA := make(map[int]float64)
	var currFleetETA float64
	var nFleet int

	for i, e := range position {
		ETA[e] = float64(target-position[i]) / float64(speed[i])
	}

	sort.Ints(position)

	for i := len(position) - 1; i >= 0; i-- {
		if ETA[position[i]] > currFleetETA {
			nFleet++
			currFleetETA = ETA[position[i]]
		}
	}

	return nFleet
}
