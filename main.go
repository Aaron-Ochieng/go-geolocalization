package main

import (
	"fmt"
	"groupie-tracker/utils"
)

func main() {
	apiIndex := utils.GetApiIndex()

	Artists := utils.GetArtists(apiIndex.Artists)
	Locations, err := utils.GetLocations(apiIndex.Locations)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Artists[0].CreationDate)
	fmt.Println(Locations[0])
}
