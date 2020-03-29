package statusprint

import (
	"fmt"
)


func Pak() bool{
	pk := true
	fmt.Println("----PAKISTAN----")
	fmt.Println("Total Cases: 1526")
	fmt.Println("Total Deaths: 13")
	return pk
}
func SK() bool{
	sk := true
	fmt.Println("----SouthKorea----")
	fmt.Println("Total Cases: 6,284")
	fmt.Println("Total Deaths: 42")
	return sk
}
func France() bool{
	fr := true
	fmt.Println("----France----")
	fmt.Println("Total Cases: 37,575")
	fmt.Println("Total Deaths: 2,314")
	return fr
}
func PerM() bool{
	pm := true
	fmt.Println("----Message----")
	fmt.Println("Stay Inside. Stay Safe.")
	return pm

}
