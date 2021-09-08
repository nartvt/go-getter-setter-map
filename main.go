package main

import "fmt"

func main() {
	mapData:=make(DataMap)
	mapData["1"] = "B"
	mapData["2"] = "C"
	mapData["3"] = "D"
	mapData["4"] = "E"

	fmt.Println(mapData)

	fmt.Println(mapData.Get("1"))

	mapData.Put("5","F")
	fmt.Println(mapData.Get("5"))
	fmt.Println(mapData)

	mapData.Delete("3")
	fmt.Println(mapData.Get("3"))
	fmt.Println(mapData)
}
