/*
Create a map called stateCapital. It will have the name of an Indian state as key and its capital as values. Initialize the same with 10 Indian state’s
capital data.
For example stateCapital[“Gujarat”] = “Gandhinagar”
Print the data type of stateCapital with %T, print stateCapital Also iterate over map and print state name and its capital name one by one,
in the following format
“Gujarat” → “Gandhinagar”
*/

package main

import "fmt"

func main() {
	stateCapital := map[string]string{
		"Gujarat":        "Gandhinagar",
		"Tamil Nadu":     "Chennai",
		"Andhra Pradesh": "Amaravati",
		"Maharashtra":    "Mumbai",
		"Karnataka":      "Bengaluru",
	}

	fmt.Println("stateCapital:", stateCapital)

	for state, capital := range stateCapital {
		fmt.Printf("%s -> %s\n", state, capital)
	}
}
