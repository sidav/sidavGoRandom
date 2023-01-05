package main

import (
	"fmt"
	prng "github.com/sidav/sidavgorandom/prng"
	"github.com/sidav/sidavgorandom/prng/fibrandom"
	"github.com/sidav/sidavgorandom/prng/pcgrandom"
)

func main() {
	mod := 2
	printNumbers := true
	const tries = 1000000000
	fmt.Println("Testing FibRandom")
	testPRNG(mod, tries, printNumbers, fibrandom.New())
	fmt.Println("Testing PCG")
	testPRNG(mod, tries, printNumbers, pcgrandom.New(-1))
}

func testPRNG(mod int, totalToGenerate int, printNumbers bool, rnd prng.PRNG) {
	occurencies := make([]int, mod)
	var maxRepeatCount, currentRepeatCount, prevForRepeatCount int

	for i := 0; i < totalToGenerate; i++ {
		r := rnd.Rand(mod)
		if r == prevForRepeatCount {
			currentRepeatCount++
		} else {
			if currentRepeatCount > maxRepeatCount {
				maxRepeatCount = currentRepeatCount
			}
			prevForRepeatCount = r
			currentRepeatCount = 0
		}

		occurencies[r]++
	}
	meanNumberExpected := totalToGenerate / mod
	min := totalToGenerate + 1
	max := -1
	for i, v := range occurencies {
		if printNumbers {
			if v >= meanNumberExpected {
				fmt.Printf("%d: %d (+%d); ", i, v, v-meanNumberExpected)
			} else {
				fmt.Printf("%d: %d (-%d); ", i, v, meanNumberExpected-v)
			}
		}
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	fmt.Println()
	fmt.Printf("Min-max diff is %d\n", max-min)
	fmt.Printf("Max repeat count was %d\n", maxRepeatCount)
	fmt.Println()
}
