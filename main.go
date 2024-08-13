package main

import (
	"fmt"
	"groupie-tracker/utils"
)

func main() {
	apiIndex := utils.GetApiIndex()

	Artists := utils.GetArtists(apiIndex.Artists)
	fmt.Println(Artists[0].CreationDate)
}
