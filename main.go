package main

import (
	"math"
)

// Return -1 if not found
func MinDistance1(arr []string, a string, b string) int {
	acc := make(map[string][]int)
	for i, key := range arr {
		if _, ok := acc[key]; !ok {
			acc[key] = []int{i + 1}
			continue
		}
		acc[key] = append(acc[key], i+1)
	}
	if _, ok := acc[a]; !ok {
		return -1
	}
	if _, ok := acc[b]; !ok {
		return -1
	}

	minDist := len(arr)

	for _, v := range acc[a] {
		for _, w := range acc[b] {
			val := int(math.Abs(float64(w - v)))
			if val < minDist {
				minDist = val
			}
		}
	}

	return minDist

}

// Return -1 if not found
func GetMin(arr []string, a, b string) int {
	var aIndex, bIndex int

	length := len(arr)
	minDist := length
	for aIndex < length && bIndex < length {
		if (arr[aIndex] == a && arr[bIndex] == b) ||
			(arr[aIndex] == b && arr[bIndex] == a) {

			val := int(math.Abs(float64(bIndex - aIndex)))
			// increase aIndex and bIndex if match condition
			aIndex += 1
			bIndex += 1
			if val < minDist {
				minDist = val
			}
			continue
		}

		// with case a appear more 1 time
		if arr[bIndex] == a {
			aIndex = bIndex
		}

		// Case when b appear before a
		if arr[aIndex] == a || arr[aIndex] == b {
			bIndex += 1
			continue
		}

		if arr[bIndex] == b {
			aIndex += 1
			continue
		}

		// if not match then increase aIndex and bIndex
		aIndex += 1
		bIndex += 1
	}
	return minDist

}

// Return -1 if not found
func MinDistance2(arr []string, a, b string) int {
	acc := make(map[string]int)
	for _, key := range arr {
		if _, ok := acc[key]; !ok {
			acc[key] = 1
			continue
		}
		acc[key] += 1
	}
	if _, ok := acc[a]; !ok {
		return -1
	}
	if _, ok := acc[b]; !ok {
		return -1
	}

	minA := GetMin(arr, a, b)
	// swap a and b
	minB := GetMin(arr, b, a)

	if minA < minB {
		return minA
	}
	return minB
}
