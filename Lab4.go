package main

import (
	sm "PP_LAB4/SomeUtils"
	mp "PP_LAB4/WorkWithMaps"
	"fmt"
)

func main() {
	m := mp.CreateMap()
	m = mp.AddNew(m)
	mp.AverAge(m)
	mp.Del(m)

	fmt.Println()
	sm.Upper()
	sm.Sum()
	sm.ReverseArr()
	fmt.Println()

}
