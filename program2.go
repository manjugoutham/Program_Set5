/*Enhance the above program to print the values sorted with keys i.e. by state names.
Hint: Create a slice of states, use sort package of go and sort the created slice
of states. Iterate over the sorted slice of states and print the state names.*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	stateCap := map[string]string{
		"Gujarat":        "Gandhinagar",
		"Tamil Nadu":     "Chennai",
		"Andhra Pradesh": "Amaravati",
		"Maharashtra":    "Mumbai",
		"Karnataka":      "Bengaluru"}
	for key, values := range stateCap {
		fmt.Printf(key, values)
	}

	for key, values := range stateCap {
		fmt.Println(key, values)
	}

	keys := make([]string, 0, len(stateCap))
	for key := range stateCap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return true
	})
	for _, l := range keys {
		fmt.Println(stateCap[l])
	}
}
