package main

import "fmt"

func main() {
	ids := []int{33, 23, 12, 45, 93, 54}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id

	}
	fmt.Println("Sum", sum)
}
