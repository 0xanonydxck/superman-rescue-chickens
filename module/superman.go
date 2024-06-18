package module

import (
	"sync"
)

/*
Questions:
	- In a one-dimensional world, superman needs to protect chickens from a heavy rainstorm using a roof of limited length.
	- Given the positions of chickens and the length of the roof Superman can carry, determine the maximum number of chickens Superman can protect.

Inputs:
	1. n (1 <= integer <= 1_000_000) - number of chickens, which superman can protect.
	2. k (1 <= integer <= 1_000_000) - denotes the length of the roof Superman can carry.
	3. p (1 <= integer array length <= 1_000_000_000) - representing the positions of the chickens on the one-dimension.

Returns:
	- Output a single integer, denotes the maximum number of chickens Superman can protect with the given roof length.
*/

func SupermanRescueChicken(protectLimit, roofLength uint, chickenPositions []uint) (protectNum uint) {
	wg := sync.WaitGroup{}
	resultCh := make(chan uint)

	for i := 0; i < len(chickenPositions)-1; i++ {
		clone := chickenPositions[i:]

		wg.Add(1)
		go Rescue(&wg, chickenPositions[i], protectLimit, roofLength, clone, resultCh)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		if protectNum < result {
			protectNum = result
		}
	}

	return
}

func Rescue(wg *sync.WaitGroup, position, protectLimit, roofLength uint, chickenPositions []uint, resultCh chan<- uint) {
	defer wg.Done()
	var roofLimit uint = roofLength + position
	var protectNum uint = 0

	for _, chickenPosition := range chickenPositions {
		if chickenPosition < roofLimit {
			protectNum += 1
		}
	}

	resultCh <- protectNum
}
