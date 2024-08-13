package main

import (
	"fmt"
	"groupie-tracker/utils"
)

func main() {
	apiIndex := utils.GetApiIndex()

	fmt.Printf("Artists API: %s\n", apiIndex.Artists)
	fmt.Printf("Locations API: %s\n", apiIndex.Locations)
	fmt.Printf("Dates API: %s\n", apiIndex.Dates)
	fmt.Printf("Relation API: %s\n", apiIndex.Relation)
}
