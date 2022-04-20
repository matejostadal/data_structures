/*
MATEJ OSTADAL
KMI UPOL
2022
*/

package main

import (
	"dastr/binomialheaps"
	"fmt"
	"sort"
)

func main() {

	// array of values that will be added to the heap
	input := []int{20, 50, 11, 7, 45, 12, 9, 3, 13, 15, 10, 5, 6, 55}

	binoheapsValuesTest(input) // testing value handling

	binoheapPrintTest(input) // testing the appearance

}

// inserts the values into an empty heap and prints it
func binoheapPrintTest(keys []int) {

	h := binomialheaps.MakeBinoHeap()

	for _, i := range keys {

		node := binomialheaps.MakeBinoNode(i)
		h.InsertNode(node)
	}

	h.PrintHeap()

}

// inserts the values into an empty heap and checks their correct order when extracting
func binoheapsValuesTest(keys []int) {

	h := binomialheaps.MakeBinoHeap()

	for _, i := range keys {

		node := binomialheaps.MakeBinoNode(i)
		h.InsertNode(node)

		h.DecreaseKey(node, node.Key-1)
	}

	// correcting values in an array (because they were decreased in the heap)
	for index, value := range keys {
		keys[index] = value - 1
	}

	// sorting the array of test values
	sort.Ints(keys)

	// this tests if the minimum of the heap is the same as the minimum of the array
	for _, number := range keys {

		min_node := h.ExtractMin()

		// if not
		if min_node != nil && number != min_node.Key {
			fmt.Printf("Failed because should be: %d, and is: %d \n", number, min_node.Key)
		}
		// if yes
		if min_node != nil && number == min_node.Key {
			fmt.Printf("Correct: %d = %d \n", number, min_node.Key)
		}
	}
	fmt.Printf("Finished!\n\n")

}
