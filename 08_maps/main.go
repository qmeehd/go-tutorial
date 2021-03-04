package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key values
	emails["Jordan"] = "jordan@gmail.com"
	emails["Emily"] = "emily@gmail.com"
	emails["Damian"] = "damian@gmail.com"

	fmt.Println(emails)
	//fmt.Println(len(emails))

	//Delete from map
	delete(emails, "Damian")
	fmt.Println(emails)
}
